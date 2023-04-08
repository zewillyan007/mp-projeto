package usecase

import (
	port_shared "mp-projeto/shared/port"
	"mp-projeto/shark/core/port"
)

type SharkChipUseCaseRemoveByIdShark struct {
	repository port.SharkChipIRepository
}

func NewSharkChipUseCaseRemoveByIdShark(repository port.SharkChipIRepository) *SharkChipUseCaseRemoveByIdShark {
	return &SharkChipUseCaseRemoveByIdShark{repository: repository}
}

func (o *SharkChipUseCaseRemoveByIdShark) WithTransaction(transaction port_shared.ITransaction) *SharkChipUseCaseRemoveByIdShark {
	return NewSharkChipUseCaseRemoveByIdShark(o.repository.WithTransaction(transaction))
}

func (o *SharkChipUseCaseRemoveByIdShark) Execute(IdShark int64) error {
	return o.repository.SqlQuery("DELETE FROM shark.shark_chip WHERE id_shark = ?", IdShark)
}
