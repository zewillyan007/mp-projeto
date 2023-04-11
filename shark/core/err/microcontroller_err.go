package err

import "errors"

var (
	MicrocontrollerErrorIdLocation   = errors.New("IdLocation cannot be null")
	MicrocontrollerErrorSerialNumber = errors.New("SerialNumber cannot be null")
	MicrocontrollerErrorModel        = errors.New("Model cannot be null")
	MicrocontrollerErrorStatus       = errors.New("Status cannot be null")
)
