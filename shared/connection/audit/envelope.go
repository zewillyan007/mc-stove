package audit

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"reflect"
	"time"
)

type Envelope struct {
	Id     uint64
	record *Record
}

func NewEnvelope(name, sender string, event EventType) *Envelope {
	envelope := &Envelope{
		record: &Record{
			Hash: "",
			Options: &Options{
				UserId:        0,
				UserName:      "",
				RecordId:      0,
				AccessAddress: nil,
				StartDateTime: 0,
				EndDateTime:   0,
				Error:         "",
			},
			Status:  StatusTypeStr[Running],
			Service: sender,
			Event:   EventTypeStr[event],
			Topic:   name,
			Tasks:   make([]*Task, 0),
		},
	}

	envelope.Id = 0

	return envelope
}

func (e *Envelope) AddTask(action EventType, entity string, recordId uint64, sqlText string, dataStruct interface{}) *Task {

	task := &Task{
		Action:     EventTypeStr[action],
		Entity:     entity,
		RecordId:   recordId,
		SqlText:    sqlText,
		TaskId:     uint32(len(e.record.Tasks) + 1),
		Data:       nil,
		dataStruct: reflect.ValueOf(dataStruct),
	}

	e.record.Tasks = append(e.record.Tasks, task)
	return task
}

func (e *Envelope) Prepare() (err error) {
	e.record.Options.UpdateThis = e.record.Hash

	if e.record.Options.StartDateTime == 0 {
		e.record.Options.StartDateTime = time.Now().Unix()
	}

	for _, t := range e.record.Tasks {
		if t.Data, err = ToData(t.dataStruct.Interface(), "json"); err != nil {
			return err
		}
		if t.RecordId == 0 {
			if id, ok := t.Data["id"]; ok {
				t.RecordId = uint64(id.(int64))
			}
		}
	}

	e.record.Hash, err = e.makeHash()
	return err
}

func (e *Envelope) SetEndTime(v time.Time) {
	e.record.Options.EndDateTime = v.Unix()
}

func (e *Envelope) SetStartTime(v time.Time) {

	e.record.Options.StartDateTime = v.Unix()
}

func (e *Envelope) SetStatus(value StatusType) {
	e.record.Status = StatusTypeStr[value]
}

func (e *Envelope) Options() *Options {
	return e.record.Options
}

func (e *Envelope) makeHash() (string, error) {

	b, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	h := sha256.New()
	h.Write(b)

	return hex.EncodeToString(h.Sum(nil)), err

}

type Task struct {
	Action     string `json:"action"`
	Entity     string `json:"entity"`
	RecordId   uint64 `json:"id_record"`
	SqlText    string `json:"sql_text"`
	TaskId     uint32 `json:"id_task"`
	Data       Data   `json:"data"`
	dataStruct reflect.Value
}

type Record struct {
	Options *Options `json:"options"`
	Status  string   `json:"status"`
	Service string   `json:"service"`
	Event   string   `json:"event"`
	Topic   string   `json:"topic"`
	Hash    string   `json:"hash"`
	Tasks   []*Task  `json:"data"`
}

type Options struct {
	Id            uint64 `json:"id"`
	UserId        uint64 `json:"id_user"`
	UserName      string `json:"user_name"`
	RecordId      uint64 `json:"id_record"`
	AccessAddress Data   `json:"access_address,omitempty"`
	StartDateTime int64  `json:"start_date_time"`
	EndDateTime   int64  `json:"end_date_time"`
	Error         string `json:"error"`
	UpdateThis    string `json:"update_this"`
}

type Input interface{}

type Data map[string]interface{}

type EventType uint8

type StatusType uint8

const (
	Insert EventType = iota
	Update
	Delete
	Read
	Unknown
)

const (
	Running StatusType = iota
	Success
	Waiting
	Error
)

var EventTypeStr = map[EventType]string{
	Insert:  "I",
	Update:  "U",
	Delete:  "D",
	Read:    "R",
	Unknown: "-",
}

var StatusTypeStr = map[StatusType]string{
	Running: "RUNNING",
	Success: "SUCCESS",
	Waiting: "WAITING",
	Error:   "ERROR",
}

const (
	offset64 = 14695981039346656037
	prime64  = 1099511628211
)

func makeId(key string) uint64 {
	var hash uint64 = offset64
	for i := 0; i < len(key); i++ {
		hash ^= uint64(key[i])
		hash *= prime64
	}
	return hash
}
