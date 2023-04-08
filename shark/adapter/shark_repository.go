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

type SharkRepository struct {
	*adapter.Repository
}

func NewSharkRepository(db *gorm.DB) *SharkRepository {

	return &SharkRepository{
		Repository: &adapter.Repository{Db: db, Table: "shark.shark"},
	}
}

func (re *SharkRepository) WithTransaction(transaction port_shared.ITransaction) port.SharkIRepository {
	return NewSharkRepository(transaction.GetTransaction())
}

func (re *SharkRepository) create(Shark *entity.Shark) (interface{}, error) {

	result := re.Insert(Shark)
	if result != nil {
		return nil, result
	}
	return Shark, nil
}

func (re *SharkRepository) update(Shark *entity.Shark) (interface{}, error) {

	err := re.Update(Shark)
	if err != nil {
		return nil, err
	}
	return Shark, nil
}

func (re *SharkRepository) Get(id int64) (*entity.Shark, error) {

	Shark := service.FactoryShark()
	re.First(Shark, id)
	return Shark, nil
}

func (re *SharkRepository) GetAll(conditions ...interface{}) []*entity.Shark {

	var Sharks []*entity.Shark
	re.Find(&Sharks, conditions...)
	return Sharks
}

func (re *SharkRepository) Save(Shark *entity.Shark) (*entity.Shark, error) {

	if Shark.Id == 0 {
		_, err := re.create(Shark)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := re.update(Shark)
		if err != nil {
			return nil, err
		}
	}
	return Shark, nil
}

func (re *SharkRepository) Remove(Shark *entity.Shark) error {

	result := re.Delete(Shark)
	if result != nil {
		return result
	}
	return nil
}

func (re *SharkRepository) SqlQuery(sql string, values ...interface{}) error {
	return re.Exec(sql, values...)
}

func (re *SharkRepository) SqlQueryRow(sql string) *sql.Row {
	return re.QueryRow(sql)
}

func (re *SharkRepository) SqlQueryRows(sql string) (*sql.Rows, error) {
	return re.QueryRows(sql)
}

func (re *SharkRepository) SqlQueryPaginator(columns string, table string, where string, sqlTemplate string, page float64, limit float64, order ...string) (map[string]interface{}, error) {
	return re.QueryPaginator(columns, table, where, sqlTemplate, page, limit, order...)
}
