package dto

type StoveDtoIn struct {
	Id               int64
	Number           int64
	Length           float64
	Width            float64
	Height           float64
	CreationDateTime string
	ChangeDateTime   string
}

func NewStoveDtoIn() *StoveDtoIn {

	return &StoveDtoIn{
		Id:               0,
		Number:           0,
		Length:           0,
		Width:            0,
		Height:           0,
		CreationDateTime: "",
		ChangeDateTime:   "",
	}
}

type StoveDtoOut struct {
	Id               int64
	Number           int64
	Length           float64
	Width            float64
	Height           float64
	CreationDateTime string
	ChangeDateTime   string
}

func NewStoveDtoOut() *StoveDtoOut {

	return &StoveDtoOut{
		Id:               0,
		Number:           0,
		Length:           0,
		Width:            0,
		Height:           0,
		CreationDateTime: "",
		ChangeDateTime:   "",
	}
}
