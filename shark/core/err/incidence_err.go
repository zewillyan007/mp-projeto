package err

import "errors"

var (
	IncidenceErrorChipNumber                  = errors.New("ChipNumber cannot be null")
	IncidenceErrorMicrocontrollerSerialNumber = errors.New("MicrocontrollerSerialNumber cannot be null")
)
