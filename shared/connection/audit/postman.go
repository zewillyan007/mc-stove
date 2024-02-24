package audit

import (
	"encoding/json"
	"strings"
)

type Postman struct {
	config    *Config
	publisher *Publisher
}

type Config struct {
	Sender    string
	Exchange  string
	QueueName string
	Publisher *PublisherConfig
}

type PublisherConfig struct {
	User,
	Password,
	Host,
	Port string
}

const envMax = 2048

func NewPostman(config *Config) (*Postman, error) {
	if publisher, err := NewPublisher(
		config.Publisher.User,
		config.Publisher.Password,
		config.Publisher.Host+":"+config.Publisher.Port,
	); err != nil {
		return nil, err
	} else {
		return &Postman{
			config:    config,
			publisher: publisher,
		}, err
	}
}

func NewDefaultPostman(config *Config) (*Postman, error) {
	var err error
	DefaultPostman, err = NewPostman(config)
	return DefaultPostman, err
}

func (p *Postman) NewEnvelope(name string, event EventType) *Envelope {
	return NewEnvelope(name, p.config.Sender, event)
}

func (p *Postman) Push(envelope *Envelope) error {
	if err := envelope.Prepare(); err != nil {
		return err
	}

	if data, err := json.Marshal(envelope.record); err != nil {
		return err
	} else {
		return p.publisher.Produce(p.config.Exchange, p.config.QueueName, "", data)
	}
}

func (p *Postman) ExtractEvent(stmt string) EventType {
	count := 0
	token := ""

	for _, i := range stmt {
		if i > ' ' {
			token += string(i)
			count = 0
		} else {
			count++
		}

		if count >= 1 && len(token) > 1 {
			break
		}
	}

	switch strings.ToLower(token) {
	case tokenInsert:
		return Insert
	case tokenUpdate:
		return Update
	case tokenDelete:
		return Delete
	case tokenRead:
		return Read
	default:
		return Unknown
	}
}

var DefaultPostman *Postman

const (
	tokenInsert = "insert"
	tokenUpdate = "update"
	tokenDelete = "delete"
	tokenRead   = "select"
)
