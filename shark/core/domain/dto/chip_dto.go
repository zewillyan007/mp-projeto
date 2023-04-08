package dto

type ChipDtoIn struct {
	Id               string
	Number           string
	Status           string
	CreationDateTime string
	ChangeDateTime   string
}

func NewChipDtoIn() *ChipDtoIn {
	return &ChipDtoIn{
		Id:               "",
		Number:           "",
		Status:           "",
		CreationDateTime: "",
		ChangeDateTime:   "",
	}
}

type ChipDtoOut struct {
	Id               string
	Number           string
	Status           string
	CreationDateTime string
	ChangeDateTime   string
}

func NewChipDtoOut() *ChipDtoOut {
	return &ChipDtoOut{
		Id:               "",
		Number:           "",
		Status:           "",
		CreationDateTime: "",
		ChangeDateTime:   "",
	}
}
