package entity

import (
	"mc-stove/plantation/core/err"
	"strings"
	"time"
)

type Plant struct {
	Id               int64      `json:"id"`
	Species          string     `json:"species"`
	Temperature      float64    `json:"length"`
	Ph               float64    `json:"ph"`
	Moisture         float64    `json:"moisture"`
	CreationDateTime *time.Time `json:"creation_date_time"`
	ChangeDateTime   *time.Time `json:"change_date_time"`
}

func NewPlant() *Plant {
	return &Plant{
		Id:               0,
		Species:          "",
		Temperature:      0,
		Ph:               0,
		Moisture:         0,
		CreationDateTime: &time.Time{},
		ChangeDateTime:   &time.Time{},
	}
}

func (ent *Plant) GetId() int64 {
	return ent.Id
}

func (ent *Plant) SetId(id int64) {
	ent.Id = id
}

func (ent *Plant) IsValid() error {
	if len(strings.TrimSpace(ent.Species)) == 0 {
		return err.PlantErrorSpecies
	}
	return nil
}
