package service

import (
	"mp-projeto/shark/core/domain/entity"
)

func FactoryIncidence() *entity.Incidence {
	return entity.NewIncidence()
}
