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

type IncidenceRepository struct {
	*adapter.Repository
}

func NewIncidenceRepository(db *gorm.DB) *IncidenceRepository {

	return &IncidenceRepository{
		Repository: &adapter.Repository{Db: db, Table: "shark.incidence"},
	}
}

func (re *IncidenceRepository) WithTransaction(transaction port_shared.ITransaction) port.IncidenceIRepository {
	return NewIncidenceRepository(transaction.GetTransaction())
}

func (re *IncidenceRepository) create(Incidence *entity.Incidence) (interface{}, error) {

	result := re.Insert(Incidence)
	if result != nil {
		return nil, result
	}
	return Incidence, nil
}

func (re *IncidenceRepository) update(Incidence *entity.Incidence) (interface{}, error) {

	err := re.Update(Incidence)
	if err != nil {
		return nil, err
	}
	return Incidence, nil
}

func (re *IncidenceRepository) Get(id int64) (*entity.Incidence, error) {

	Incidence := service.FactoryIncidence()
	re.First(Incidence, id)
	return Incidence, nil
}

func (re *IncidenceRepository) GetAll(conditions ...interface{}) []*entity.Incidence {

	var Incidences []*entity.Incidence
	re.Find(&Incidences, conditions...)
	return Incidences
}

func (re *IncidenceRepository) Save(Incidence *entity.Incidence) (*entity.Incidence, error) {

	if Incidence.Id == 0 {
		_, err := re.create(Incidence)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := re.update(Incidence)
		if err != nil {
			return nil, err
		}
	}
	return Incidence, nil
}

func (re *IncidenceRepository) Remove(Incidence *entity.Incidence) error {

	result := re.Delete(Incidence)
	if result != nil {
		return result
	}
	return nil
}

func (re *IncidenceRepository) SqlQuery(sql string, values ...interface{}) error {
	return re.Exec(sql, values...)
}

func (re *IncidenceRepository) SqlQueryRow(sql string) *sql.Row {
	return re.QueryRow(sql)
}

func (re *IncidenceRepository) SqlQueryRows(sql string) (*sql.Rows, error) {
	return re.QueryRows(sql)
}

func (re *IncidenceRepository) SqlQueryPaginator(columns string, table string, where string, sqlTemplate string, page float64, limit float64, order ...string) (map[string]interface{}, error) {
	return re.QueryPaginator(columns, table, where, sqlTemplate, page, limit, order...)
}
