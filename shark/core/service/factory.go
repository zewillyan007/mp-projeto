package service

import (
	"mp-projeto/shark/core/domain/entity"
)

func FactoryIncidence() *entity.Incidence {
	return entity.NewIncidence()
}

func FactoryChip() *entity.Chip {
	return entity.NewChip()
}

func FactoryShark() *entity.Shark {
	return entity.NewShark()
}

func FactorySharkChip() *entity.SharkChip {
	return entity.NewSharkChip()
}

func FactorySex() *entity.Sex {
	return entity.NewSex()
}

func FactoryChipStatusType() *entity.ChipStatusType {
	return entity.NewChipStatusType()
}

func FactorySharkChipStatusType() *entity.SharkChipStatusType {
	return entity.NewSharkChipStatusType()
}
