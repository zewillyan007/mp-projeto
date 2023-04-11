package usecase

import (
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type MicrocontrollerUseCaseGet struct {
	repository port.MicrocontrollerIRepository
}

func NewMicrocontrollerUseCaseGet(repository port.MicrocontrollerIRepository) *MicrocontrollerUseCaseGet {
	return &MicrocontrollerUseCaseGet{repository: repository}
}

func (o *MicrocontrollerUseCaseGet) Execute(id int64) (*entity.Microcontroller, error) {

	entity, err := o.repository.Get(id)

	if err != nil {
		return nil, err
	}
	return entity, nil
}
