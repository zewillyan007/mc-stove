package dto

type TypeMicrocontrollerDtoIn struct {
	Id               int64
	Name             string
	Mnemonic         string
	Hint             string
	CreationDateTime string
	ChangeDateTime   string
}

func NewTypeMicrocontrollerDtoIn() *TypeMicrocontrollerDtoIn {

	return &TypeMicrocontrollerDtoIn{
		Id:               0,
		Name:             "",
		Mnemonic:         "",
		Hint:             "",
		CreationDateTime: "",
		ChangeDateTime:   "",
	}
}

type TypeMicrocontrollerDtoOut struct {
	Id               int64
	Name             string
	Mnemonic         string
	Hint             string
	CreationDateTime string
	ChangeDateTime   string
}

func NewTypeMicrocontrollerDtoOut() *TypeMicrocontrollerDtoOut {

	return &TypeMicrocontrollerDtoOut{
		Id:               0,
		Name:             "",
		Mnemonic:         "",
		Hint:             "",
		CreationDateTime: "",
		ChangeDateTime:   "",
	}
}
