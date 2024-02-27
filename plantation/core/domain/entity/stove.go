package entity

import (
	"mc-stove/plantation/core/err"
	"time"
)

type Stove struct {
	Id               int64      `json:"id"`
	Number           int64      `json:"number"`
	Length           float64    `json:"length"`
	Width            float64    `json:"width"`
	Height           float64    `json:"height"`
	CreationDateTime *time.Time `json:"creation_date_time"`
	ChangeDateTime   *time.Time `json:"change_date_time"`
}

func NewStove() *Stove {
	return &Stove{
		Id:               0,
		Number:           0,
		Length:           0,
		Width:            0,
		Height:           0,
		CreationDateTime: &time.Time{},
		ChangeDateTime:   &time.Time{},
	}
}

func (ent *Stove) GetId() int64 {
	return ent.Id
}

func (ent *Stove) SetId(id int64) {
	ent.Id = id
}

func (ent *Stove) IsValid() error {
	if ent.Length == 0 {
		return err.StoveErrorLength
	}
	if ent.Width == 0 {
		return err.StoveErrorWidth
	}
	if ent.Height == 0 {
		return err.StoveErrorHeight
	}
	if ent.Number == 0 {
		return err.StoveErrorNumber
	}
	return nil
}
