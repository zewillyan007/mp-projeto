package dto

type SharkChipStatusTypeDtoIn struct {
	Id               string
	Name             string
	Mnemonic         string
	Hint             string
	CreationDateTime string
	ChangeDateTime   string
	DisableDateTime  string
}

func NewSharkChipStatusTypeDtoIn() *SharkChipStatusTypeDtoIn {

	return &SharkChipStatusTypeDtoIn{
		Id:               "",
		Name:             "",
		Mnemonic:         "",
		Hint:             "",
		CreationDateTime: "",
		ChangeDateTime:   "",
		DisableDateTime:  "",
	}
}

type SharkChipStatusTypeDtoOut struct {
	Id               string
	Name             string
	Mnemonic         string
	Hint             string
	CreationDateTime string
	ChangeDateTime   string
	DisableDateTime  string
}

func NewSharkChipStatusTypeDtoOut() *SharkChipStatusTypeDtoOut {

	return &SharkChipStatusTypeDtoOut{
		Id:               "",
		Name:             "",
		Mnemonic:         "",
		Hint:             "",
		CreationDateTime: "",
		ChangeDateTime:   "",
		DisableDateTime:  "",
	}
}
