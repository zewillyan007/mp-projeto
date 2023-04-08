package usecase

import (
	port_shared "mp-projeto/shared/port"
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
)

type SharkUseCaseSave struct {
	repository port.SharkIRepository
}

func NewSharkUseCaseSave(repository port.SharkIRepository) *SharkUseCaseSave {
	return &SharkUseCaseSave{repository: repository}
}

func (o *SharkUseCaseSave) WithTransaction(transaction port_shared.ITransaction) *SharkUseCaseSave {
	return NewSharkUseCaseSave(o.repository.WithTransaction(transaction))
}

func (o *SharkUseCaseSave) Execute(Shark *entity.Shark) (*entity.Shark, error) {

	err := Shark.IsValid()

	if err != nil {
		return nil, err
	}

	entity, err := o.repository.Save(Shark)

	if err != nil {
		return nil, err
	}

	return entity, nil
}
