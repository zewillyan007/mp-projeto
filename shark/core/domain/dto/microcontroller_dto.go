package dto

type MicrocontrollerDtoIn struct {
	Id               string
	IdLocation       string
	SerialNumber     string
	Model            string
	Status           string
	CreationDateTime string
	ChangeDateTime   string
}

func NewMicrocontrollerDtoIn() *MicrocontrollerDtoIn {
	return &MicrocontrollerDtoIn{
		Id:               "",
		IdLocation:       "",
		SerialNumber:     "",
		Model:            "",
		Status:           "",
		CreationDateTime: "",
		ChangeDateTime:   "",
	}
}

type MicrocontrollerDtoOut struct {
	Id               string
	IdLocation       string
	SerialNumber     string
	Model            string
	Status           string
	CreationDateTime string
	ChangeDateTime   string
}

func NewMicrocontrollerDtoOut() *MicrocontrollerDtoOut {
	return &MicrocontrollerDtoOut{
		Id:               "",
		IdLocation:       "",
		SerialNumber:     "",
		Model:            "",
		Status:           "",
		CreationDateTime: "",
		ChangeDateTime:   "",
	}
}
