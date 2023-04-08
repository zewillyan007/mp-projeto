package dto

type ChipStatusTypeDtoIn struct {
	Id               string
	Name             string
	Mnemonic         string
	Hint             string
	CreationDateTime string
	ChangeDateTime   string
	DisableDateTime  string
}

func NewChipStatusTypeDtoIn() *ChipStatusTypeDtoIn {

	return &ChipStatusTypeDtoIn{
		Id:               "",
		Name:             "",
		Mnemonic:         "",
		Hint:             "",
		CreationDateTime: "",
		ChangeDateTime:   "",
		DisableDateTime:  "",
	}
}

type ChipStatusTypeDtoOut struct {
	Id               string
	Name             string
	Mnemonic         string
	Hint             string
	CreationDateTime string
	ChangeDateTime   string
	DisableDateTime  string
}

func NewChipStatusTypeDtoOut() *ChipStatusTypeDtoOut {

	return &ChipStatusTypeDtoOut{
		Id:               "",
		Name:             "",
		Mnemonic:         "",
		Hint:             "",
		CreationDateTime: "",
		ChangeDateTime:   "",
		DisableDateTime:  "",
	}
}
