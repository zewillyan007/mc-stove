package dto

type MicrocontrollerStoveDtoIn struct {
	Id                int64
	IdStove           int64
	IdMicrocontroller int64
	CreationDateTime  string
	ChangeDateTime    string
}

func NewMicrocontrollerStoveDtoIn() *MicrocontrollerStoveDtoIn {

	return &MicrocontrollerStoveDtoIn{
		Id:                0,
		IdStove:           0,
		IdMicrocontroller: 0,
		CreationDateTime:  "",
		ChangeDateTime:    "",
	}
}

type MicrocontrollerStoveDtoOut struct {
	Id                int64
	IdStove           int64
	IdMicrocontroller int64
	CreationDateTime  string
	ChangeDateTime    string
}

func NewMicrocontrollerStoveDtoOut() *MicrocontrollerStoveDtoOut {

	return &MicrocontrollerStoveDtoOut{
		Id:                0,
		IdStove:           0,
		IdMicrocontroller: 0,
		CreationDateTime:  "",
		ChangeDateTime:    "",
	}
}
