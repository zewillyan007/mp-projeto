package dto

type IncidenceDtoIn struct {
	Id                          string
	ChipNumber                  string
	MicrocontrollerSerialNumber string
	IncidenceDateTime           string
}

func NewIncidenceDtoIn() *IncidenceDtoIn {
	return &IncidenceDtoIn{
		Id:                          "",
		ChipNumber:                  "",
		MicrocontrollerSerialNumber: "",
		IncidenceDateTime:           "",
	}
}

type IncidenceDtoOut struct {
	Id                          string
	ChipNumber                  string
	MicrocontrollerSerialNumber string
	Shark                       *SharkDtoOut
	Microcontroller             *MicrocontrollerDtoOut
	Location                    *LocationDtoOut
	IncidenceDateTime           string
}

func NewIncidenceDtoOut() *IncidenceDtoOut {
	return &IncidenceDtoOut{
		Id:                          "",
		ChipNumber:                  "",
		MicrocontrollerSerialNumber: "",
		Shark:                       &SharkDtoOut{},
		Microcontroller:             &MicrocontrollerDtoOut{},
		Location:                    &LocationDtoOut{},
		IncidenceDateTime:           "",
	}
}
