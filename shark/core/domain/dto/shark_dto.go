package dto

type SharkDtoIn struct {
	Id               string
	Species          string
	Length           string
	Weight           string
	Sex              string
	CreationDateTime string
	ChangeDateTime   string
}

func NewSharkDtoIn() *SharkDtoIn {
	return &SharkDtoIn{
		Id:               "",
		Species:          "",
		Length:           "",
		Weight:           "",
		Sex:              "",
		CreationDateTime: "",
		ChangeDateTime:   "",
	}
}

type SharkDtoOut struct {
	Id               string
	Species          string
	Length           string
	Weight           string
	Sex              string
	CreationDateTime string
	ChangeDateTime   string
}

func NewSharkDtoOut() *SharkDtoOut {
	return &SharkDtoOut{
		Id:               "",
		Species:          "",
		Length:           "",
		Weight:           "",
		Sex:              "",
		CreationDateTime: "",
		ChangeDateTime:   "",
	}
}
