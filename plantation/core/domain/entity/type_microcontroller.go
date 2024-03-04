package entity

import (
	"time"
)

type TypeMicrocontroller struct {
	Id               int64      `json:"id"`
	Name             string     `json:"name"`
	Mnemonic         string     `json:"mnemonic"`
	Hint             string     `json:"hint"`
	CreationDateTime *time.Time `json:"creation_date_time"`
	ChangeDateTime   *time.Time `json:"change_date_time"`
}

func NewTypeMicrocontroller() *TypeMicrocontroller {
	return &TypeMicrocontroller{
		Id:               0,
		Name:             "",
		Mnemonic:         "",
		Hint:             "",
		CreationDateTime: &time.Time{},
		ChangeDateTime:   &time.Time{},
	}
}

func (ent *TypeMicrocontroller) GetId() int64 {
	return ent.Id
}

func (ent *TypeMicrocontroller) SetId(id int64) {
	ent.Id = id
}

func (ent *TypeMicrocontroller) IsValid() error {
	// if len(strings.TrimSpace(ent.Species)) == 0 {
	// 	return err.TypeMicrocontrollerErrorSpecies
	// }
	return nil
}
