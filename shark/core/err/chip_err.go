package err

import "errors"

var (
	ChipErrorNumber = errors.New("Number cannot be null")
	ChipErrorStatus = errors.New("Status cannot be null")
)
