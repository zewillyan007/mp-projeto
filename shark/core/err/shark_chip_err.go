package err

import "errors"

var (
	SharkChipErrorIdShark    = errors.New("IdShark cannot be null")
	SharkChipErrorIdChip     = errors.New("IdChip cannot be null")
	SharkChipErrorChipNumber = errors.New("ChipNumber cannot be null")
	SharkChipErrorStatus     = errors.New("Status cannot be null")
)
