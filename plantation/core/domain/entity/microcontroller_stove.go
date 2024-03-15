package entity

import (
	"time"
)

type MicrocontrollerStove struct {
	Id                int64      `json:"id"`
	IdStove           int64      `json:"id_stove"`
	IdMicrocontroller int64      `json:"id_plant"`
	CreationDateTime  *time.Time `json:"creation_date_time"`
	ChangeDateTime    *time.Time `json:"change_date_time"`
}

func NewMicrocontrollerStove() *MicrocontrollerStove {
	return &MicrocontrollerStove{
		Id:                0,
		IdStove:           0,
		IdMicrocontroller: 0,
		CreationDateTime:  &time.Time{},
		ChangeDateTime:    &time.Time{},
	}
}

func (ent *MicrocontrollerStove) GetId() int64 {
	return ent.Id
}

func (ent *MicrocontrollerStove) SetId(id int64) {
	ent.Id = id
}

func (ent *MicrocontrollerStove) IsValid() error {
	// if len(strings.TrimSpace(ent.Species)) == 0 {
	// 	return err.MicrocontrollerStoveErrorSpecies
	// }
	return nil
}
