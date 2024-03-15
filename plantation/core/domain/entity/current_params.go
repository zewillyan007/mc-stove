package entity

import (
	"time"
)

type CurrentParams struct {
	Id               int64      `json:"id"`
	IdStove          int64      `json:"id_stove"`
	IdPlant          int64      `json:"id_plant"`
	StoveNumber      int64      `json:"stove_number"`
	PlantSpecies     string     `json:"plant_species"`
	Temperature      float64    `json:"temperature"`
	Moisture         float64    `json:"moisture"`
	CreationDateTime *time.Time `json:"creation_date_time"`
}

func NewCurrentParams() *CurrentParams {
	return &CurrentParams{
		Id:               0,
		IdStove:          0,
		IdPlant:          0,
		StoveNumber:      0,
		PlantSpecies:     "",
		Temperature:      0,
		Moisture:         0,
		CreationDateTime: &time.Time{},
	}
}

func (ent *CurrentParams) GetId() int64 {
	return ent.Id
}

func (ent *CurrentParams) SetId(id int64) {
	ent.Id = id
}

func (ent *CurrentParams) IsValid() error {
	// if len(strings.TrimSpace(ent.Species)) == 0 {
	// 	return err.CurrentParamsErrorSpecies
	// }
	return nil
}
