package port

import (
	"database/sql"
	port_shared "mp-projeto/shared/port"

	"mp-projeto/shark/core/domain/entity"
)

type SharkIRepository interface {
	GetAll(conditions ...interface{}) []*entity.Shark
	Remove(*entity.Shark) error
	Get(int64) (*entity.Shark, error)
	Save(*entity.Shark) (*entity.Shark, error)
	SqlQuery(string, ...interface{}) error
	SqlQueryRow(string) *sql.Row
	SqlQueryRows(string) (*sql.Rows, error)
	SqlQueryPaginator(columns string, table string, where string, sqlTemplate string, page float64, limit float64, order ...string) (map[string]interface{}, error)
	WithTransaction(transaction port_shared.ITransaction) SharkIRepository
}
