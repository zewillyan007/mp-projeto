package port

import (
	port_shared "mp-projeto/shared/port"

	"mp-projeto/shark/core/domain/entity"
)

type SexIRepository interface {
	GetAll(conditions ...interface{}) []*entity.Sex
	Remove(*entity.Sex) error
	Get(int64) (*entity.Sex, error)
	WithTransaction(transaction port_shared.ITransaction) SexIRepository
}
