package port

import (
	port_shared "mp-projeto/shared/port"

	"mp-projeto/shark/core/domain/entity"
)

type ChipStatusTypeIRepository interface {
	GetAll(conditions ...interface{}) []*entity.ChipStatusType
	Remove(*entity.ChipStatusType) error
	Get(int64) (*entity.ChipStatusType, error)
	WithTransaction(transaction port_shared.ITransaction) ChipStatusTypeIRepository
}
