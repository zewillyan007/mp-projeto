package dto

type SharkChipDtoIn struct {
	Id               string
	IdShark          string
	IdChip           string
	ChipNumber       string
	Status           string
	CreationDateTime string
}

func NewSharkChipDtoIn() *SharkChipDtoIn {
	return &SharkChipDtoIn{
		Id:               "",
		IdShark:          "",
		IdChip:           "",
		ChipNumber:       "",
		Status:           "",
		CreationDateTime: "",
	}
}

type SharkChipDtoOut struct {
	Id               string
	IdShark          string
	IdChip           string
	ChipNumber       string
	Status           string
	CreationDateTime string
}

func NewSharkChipDtoOut() *SharkChipDtoOut {
	return &SharkChipDtoOut{
		Id:               "",
		IdShark:          "",
		IdChip:           "",
		ChipNumber:       "",
		Status:           "",
		CreationDateTime: "",
	}
}
