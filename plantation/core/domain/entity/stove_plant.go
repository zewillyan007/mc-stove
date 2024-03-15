package entity

import (
	"time"
)

type StovePlant struct {
	Id               int64      `json:"id"`
	IdStove          int64      `json:"id_stove"`
	IdPlant          int64      `json:"id_plant"`
	CreationDateTime *time.Time `json:"creation_date_time"`
	ChangeDateTime   *time.Time `json:"change_date_time"`
}

func NewStovePlant() *StovePlant {
	return &StovePlant{
		Id:               0,
		IdStove:          0,
		IdPlant:          0,
		CreationDateTime: &time.Time{},
		ChangeDateTime:   &time.Time{},
	}
}

func (ent *StovePlant) GetId() int64 {
	return ent.Id
}

func (ent *StovePlant) SetId(id int64) {
	ent.Id = id
}

func (ent *StovePlant) IsValid() error {
	// if len(strings.TrimSpace(ent.Species)) == 0 {
	// 	return err.StovePlantErrorSpecies
	// }
	return nil
}
