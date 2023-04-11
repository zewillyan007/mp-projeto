package port

import (
	"database/sql"
	port_shared "mp-projeto/shared/port"

	"mp-projeto/shark/core/domain/entity"
)

type LocationIRepository interface {
	GetAll(conditions ...interface{}) []*entity.Location
	Remove(*entity.Location) error
	Get(int64) (*entity.Location, error)
	Save(*entity.Location) (*entity.Location, error)
	SqlQuery(string, ...interface{}) error
	SqlQueryRow(string) *sql.Row
	SqlQueryRows(string) (*sql.Rows, error)
	SqlQueryPaginator(columns string, table string, where string, sqlTemplate string, page float64, limit float64, order ...string) (map[string]interface{}, error)
	WithTransaction(transaction port_shared.ITransaction) LocationIRepository
}
