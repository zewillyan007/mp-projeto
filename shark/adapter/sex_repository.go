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

type SexRepository struct {
	*adapter.Repository
}

func NewSexRepository(db *gorm.DB) *SexRepository {

	return &SexRepository{
		Repository: &adapter.Repository{Db: db, Table: "domain.sex"},
	}
}

func (re *SexRepository) WithTransaction(transaction port_shared.ITransaction) port.SexIRepository {
	return NewSexRepository(transaction.GetTransaction())
}

func (re *SexRepository) create(Sex *entity.Sex) (interface{}, error) {

	result := re.Insert(Sex)
	if result != nil {
		return nil, result
	}
	return Sex, nil
}

func (re *SexRepository) update(Sex *entity.Sex) (interface{}, error) {

	err := re.Update(Sex)
	if err != nil {
		return nil, err
	}
	return Sex, nil
}

func (re *SexRepository) Get(id int64) (*entity.Sex, error) {

	Sex := service.FactorySex()
	re.First(Sex, id)
	return Sex, nil
}

func (re *SexRepository) GetAll(conditions ...interface{}) []*entity.Sex {

	var Sexs []*entity.Sex
	re.Find(&Sexs, conditions...)
	return Sexs
}

func (re *SexRepository) Save(Sex *entity.Sex) (*entity.Sex, error) {

	if Sex.Id == 0 {
		_, err := re.create(Sex)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := re.update(Sex)
		if err != nil {
			return nil, err
		}
	}
	return Sex, nil
}

func (re *SexRepository) Remove(Sex *entity.Sex) error {

	result := re.Delete(Sex)
	if result != nil {
		return result
	}
	return nil
}

func (re *SexRepository) SqlQuery(sql string, values ...interface{}) error {
	return re.Exec(sql, values...)
}

func (re *SexRepository) SqlQueryRow(sql string) *sql.Row {
	return re.QueryRow(sql)
}

func (re *SexRepository) SqlQueryRows(sql string) (*sql.Rows, error) {
	return re.QueryRows(sql)
}

func (re *SexRepository) SqlQueryPaginator(columns string, table string, where string, sqlTemplate string, page float64, limit float64, order ...string) (map[string]interface{}, error) {
	return re.QueryPaginator(columns, table, where, sqlTemplate, page, limit, order...)
}
