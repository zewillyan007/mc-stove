package dto

type MicrocontrollerDtoIn struct {
	Id               int64
	Type             string
	SerialNumber     string
	CreationDateTime string
	ChangeDateTime   string
}

func NewMicrocontrollerDtoIn() *MicrocontrollerDtoIn {

	return &MicrocontrollerDtoIn{
		Id:               0,
		Type:             "",
		SerialNumber:     "",
		CreationDateTime: "",
		ChangeDateTime:   "",
	}
}

type MicrocontrollerDtoOut struct {
	Id               int64
	Type             string
	SerialNumber     string
	CreationDateTime string
	ChangeDateTime   string
}

func NewMicrocontrollerDtoOut() *MicrocontrollerDtoOut {

	return &MicrocontrollerDtoOut{
		Id:               0,
		Type:             "",
		SerialNumber:     "",
		CreationDateTime: "",
		ChangeDateTime:   "",
	}
}
