package adapter

import (
	"database/sql"
	"mp-projeto/shared/adapter"
	port_shared "mp-projeto/shared/port"
	"mp-projeto/shark/core/domain/entity"
	"mp-projeto/shark/core/port"
	"mp-projeto/shark/core/service"

	"gorm.io/gorm"
)

type LocationRepository struct {
	*adapter.Repository
}

func NewLocationRepository(db *gorm.DB) *LocationRepository {

	return &LocationRepository{
		Repository: &adapter.Repository{Db: db, Table: "shark.location"},
	}
}

func (re *LocationRepository) WithTransaction(transaction port_shared.ITransaction) port.LocationIRepository {
	return NewLocationRepository(transaction.GetTransaction())
}

func (re *LocationRepository) create(Location *entity.Location) (interface{}, error) {

	result := re.Insert(Location)
	if result != nil {
		return nil, result
	}
	return Location, nil
}

func (re *LocationRepository) update(Location *entity.Location) (interface{}, error) {

	err := re.Update(Location)
	if err != nil {
		return nil, err
	}
	return Location, nil
}

func (re *LocationRepository) Get(id int64) (*entity.Location, error) {

	Location := service.FactoryLocation()
	re.First(Location, id)
	return Location, nil
}

func (re *LocationRepository) GetAll(conditions ...interface{}) []*entity.Location {

	var Locations []*entity.Location
	re.Find(&Locations, conditions...)
	return Locations
}

func (re *LocationRepository) Save(Location *entity.Location) (*entity.Location, error) {

	if Location.Id == 0 {
		_, err := re.create(Location)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := re.update(Location)
		if err != nil {
			return nil, err
		}
	}
	return Location, nil
}

func (re *LocationRepository) Remove(Location *entity.Location) error {

	result := re.Delete(Location)
	if result != nil {
		return result
	}
	return nil
}

func (re *LocationRepository) SqlQuery(sql string, values ...interface{}) error {
	return re.Exec(sql, values...)
}

func (re *LocationRepository) SqlQueryRow(sql string) *sql.Row {
	return re.QueryRow(sql)
}

func (re *LocationRepository) SqlQueryRows(sql string) (*sql.Rows, error) {
	return re.QueryRows(sql)
}

func (re *LocationRepository) SqlQueryPaginator(columns string, table string, where string, sqlTemplate string, page float64, limit float64, order ...string) (map[string]interface{}, error) {
	return re.QueryPaginator(columns, table, where, sqlTemplate, page, limit, order...)
}
