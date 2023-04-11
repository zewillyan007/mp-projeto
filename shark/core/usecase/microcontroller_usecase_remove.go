package usecase

import (
	port_shared "mp-projeto/shared/port"
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type MicrocontrollerUseCaseRemove struct {
	repository port.MicrocontrollerIRepository
}

func NewMicrocontrollerUseCaseRemove(repository port.MicrocontrollerIRepository) *MicrocontrollerUseCaseRemove {
	return &MicrocontrollerUseCaseRemove{repository: repository}
}

func (o *MicrocontrollerUseCaseRemove) WithTransaction(transaction port_shared.ITransaction) *MicrocontrollerUseCaseRemove {
	return NewMicrocontrollerUseCaseRemove(o.repository.WithTransaction(transaction))
}

func (o *MicrocontrollerUseCaseRemove) Execute(Microcontroller *entity.Microcontroller) error {
	return o.repository.Remove(Microcontroller)
}
