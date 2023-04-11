package dto

type LocationDtoIn struct {
	Id               string
	Name             string
	CreationDateTime string
	ChangeDateTime   string
}

func NewLocationDtoIn() *LocationDtoIn {
	return &LocationDtoIn{
		Id:               "",
		Name:             "",
		CreationDateTime: "",
		ChangeDateTime:   "",
	}
}

type LocationDtoOut struct {
	Id               string
	Name             string
	CreationDateTime string
	ChangeDateTime   string
}

func NewLocationDtoOut() *LocationDtoOut {
	return &LocationDtoOut{
		Id:               "",
		Name:             "",
		CreationDateTime: "",
		ChangeDateTime:   "",
	}
}
