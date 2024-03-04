package entity

import (
	"time"
)

type Microcontroller struct {
	Id               int64      `json:"id"`
	Type             string     `json:"type"`
	SerialNumber     string     `json:"serial_number"`
	CreationDateTime *time.Time `json:"creation_date_time"`
	ChangeDateTime   *time.Time `json:"change_date_time"`
}

func NewMicrocontroller() *Microcontroller {
	return &Microcontroller{
		Id:               0,
		Type:             "",
		SerialNumber:     "",
		CreationDateTime: &time.Time{},
		ChangeDateTime:   &time.Time{},
	}
}

func (ent *Microcontroller) GetId() int64 {
	return ent.Id
}

func (ent *Microcontroller) SetId(id int64) {
	ent.Id = id
}

func (ent *Microcontroller) IsValid() error {
	// if len(strings.TrimSpace(ent.Species)) == 0 {
	// 	return err.MicrocontrollerErrorSpecies
	// }
	return nil
}
