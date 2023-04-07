package port

import (
	"database/sql"
	port_shared "mp-projeto/shared/port"

	"mp-projeto/shark/core/domain/entity"
)

type IncidenceIRepository interface {
	GetAll(conditions ...interface{}) []*entity.Incidence
	Remove(*entity.Incidence) error
	Get(int64) (*entity.Incidence, error)
	Save(*entity.Incidence) (*entity.Incidence, error)
	SqlQuery(string, ...interface{}) error
	SqlQueryRow(string) *sql.Row
	SqlQueryRows(string) (*sql.Rows, error)
	SqlQueryPaginator(columns string, table string, where string, sqlTemplate string, page float64, limit float64, order ...string) (map[string]interface{}, error)
	WithTransaction(transaction port_shared.ITransaction) IncidenceIRepository
}
