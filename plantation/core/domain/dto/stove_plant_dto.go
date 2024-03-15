package dto

type StovePlantDtoIn struct {
	Id               int64
	IdStove          int64
	IdPlant          int64
	CreationDateTime string
	ChangeDateTime   string
}

func NewStovePlantDtoIn() *StovePlantDtoIn {

	return &StovePlantDtoIn{
		Id:               0,
		IdStove:          0,
		IdPlant:          0,
		CreationDateTime: "",
		ChangeDateTime:   "",
	}
}

type StovePlantDtoOut struct {
	Id               int64
	IdStove          int64
	IdPlant          int64
	CreationDateTime string
	ChangeDateTime   string
}

func NewStovePlantDtoOut() *StovePlantDtoOut {

	return &StovePlantDtoOut{
		Id:               0,
		IdStove:          0,
		IdPlant:          0,
		CreationDateTime: "",
		ChangeDateTime:   "",
	}
}
