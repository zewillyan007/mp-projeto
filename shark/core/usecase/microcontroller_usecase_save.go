package usecase

import (
	port_shared "mp-projeto/shared/port"
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type MicrocontrollerUseCaseSave struct {
	repository port.MicrocontrollerIRepository
}

func NewMicrocontrollerUseCaseSave(repository port.MicrocontrollerIRepository) *MicrocontrollerUseCaseSave {
	return &MicrocontrollerUseCaseSave{repository: repository}
}

func (o *MicrocontrollerUseCaseSave) WithTransaction(transaction port_shared.ITransaction) *MicrocontrollerUseCaseSave {
	return NewMicrocontrollerUseCaseSave(o.repository.WithTransaction(transaction))
}

func (o *MicrocontrollerUseCaseSave) Execute(Microcontroller *entity.Microcontroller) (*entity.Microcontroller, error) {

	err := Microcontroller.IsValid()

	if err != nil {
		return nil, err
	}

	entity, err := o.repository.Save(Microcontroller)

	if err != nil {
		return nil, err
	}

	return entity, nil
}
