package dto

type CurrentParamsDtoIn struct {
	Id                int64
	MicroSerialNumber string
	Temperature       float64
	Moisture          float64
	CreationDateTime  string
}

func NewCurrentParamsDtoIn() *CurrentParamsDtoIn {

	return &CurrentParamsDtoIn{
		Id:                0,
		MicroSerialNumber: "",
		Temperature:       0,
		Moisture:          0,
		CreationDateTime:  "",
	}
}

type CurrentParamsDtoOut struct {
	Id                int64
	MicroSerialNumber string
	Temperature       float64
	Moisture          float64
	CreationDateTime  string
}

func NewCurrentParamsDtoOut() *CurrentParamsDtoOut {

	return &CurrentParamsDtoOut{
		Id:                0,
		MicroSerialNumber: "",
		Temperature:       0,
		Moisture:          0,
		CreationDateTime:  "",
	}
}
