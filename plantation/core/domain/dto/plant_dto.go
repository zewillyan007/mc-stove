package dto

type PlantDtoIn struct {
	Id               int64
	Species          string
	Temperature      float64
	Ph               float64
	Moisture         float64
	CreationDateTime string
	ChangeDateTime   string
}

func NewPlantDtoIn() *PlantDtoIn {

	return &PlantDtoIn{
		Id:               0,
		Species:          "",
		Temperature:      0,
		Ph:               0,
		Moisture:         0,
		CreationDateTime: "",
		ChangeDateTime:   "",
	}
}

type PlantDtoOut struct {
	Id               int64
	Species          string
	Temperature      float64
	Ph               float64
	Moisture         float64
	CreationDateTime string
	ChangeDateTime   string
}

func NewPlantDtoOut() *PlantDtoOut {

	return &PlantDtoOut{
		Id:               0,
		Species:          "",
		Temperature:      0,
		Ph:               0,
		Moisture:         0,
		CreationDateTime: "",
		ChangeDateTime:   "",
	}
}
