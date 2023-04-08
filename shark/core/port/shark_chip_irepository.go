package port

import (
	"database/sql"
	port_shared "mp-projeto/shared/port"

	"mp-projeto/shark/core/domain/entity"
)

type SharkChipIRepository interface {
	GetAll(conditions ...interface{}) []*entity.SharkChip
	Remove(*entity.SharkChip) error
	Get(int64) (*entity.SharkChip, error)
	Save(*entity.SharkChip) (*entity.SharkChip, error)
	SqlQuery(string, ...interface{}) error
	SqlQueryRow(string) *sql.Row
	SqlQueryRows(string) (*sql.Rows, error)
	SqlQueryPaginator(columns string, table string, where string, sqlTemplate string, page float64, limit float64, order ...string) (map[string]interface{}, error)
	WithTransaction(transaction port_shared.ITransaction) SharkChipIRepository
}
