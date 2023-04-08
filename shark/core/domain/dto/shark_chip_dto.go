package dto

type SharkChipDtoIn struct {
	Id               string
	IdShark          string
	IdChip           string
	CreationDateTime string
}

func NewSharkChipDtoIn() *SharkChipDtoIn {
	return &SharkChipDtoIn{
		Id:               "",
		IdShark:          "",
		IdChip:           "",
		CreationDateTime: "",
	}
}

type SharkChipDtoOut struct {
	Id               string
	IdShark          string
	IdChip           string
	CreationDateTime string
}

func NewSharkChipDtoOut() *SharkChipDtoOut {
	return &SharkChipDtoOut{
		Id:               "",
		IdShark:          "",
		IdChip:           "",
		CreationDateTime: "",
	}
}
