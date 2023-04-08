package port

import (
	port_shared "mp-projeto/shared/port"

	"mp-projeto/shark/core/domain/entity"
)

type SharkChipStatusTypeIRepository interface {
	GetAll(conditions ...interface{}) []*entity.SharkChipStatusType
	Remove(*entity.SharkChipStatusType) error
	Get(int64) (*entity.SharkChipStatusType, error)
	WithTransaction(transaction port_shared.ITransaction) SharkChipStatusTypeIRepository
}
