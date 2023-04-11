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

type MicrocontrollerRepository struct {
	*adapter.Repository
}

func NewMicrocontrollerRepository(db *gorm.DB) *MicrocontrollerRepository {

	return &MicrocontrollerRepository{
		Repository: &adapter.Repository{Db: db, Table: "shark.microcontroller"},
	}
}

func (re *MicrocontrollerRepository) WithTransaction(transaction port_shared.ITransaction) port.MicrocontrollerIRepository {
	return NewMicrocontrollerRepository(transaction.GetTransaction())
}

func (re *MicrocontrollerRepository) create(Microcontroller *entity.Microcontroller) (interface{}, error) {

	result := re.Insert(Microcontroller)
	if result != nil {
		return nil, result
	}
	return Microcontroller, nil
}

func (re *MicrocontrollerRepository) update(Microcontroller *entity.Microcontroller) (interface{}, error) {

	err := re.Update(Microcontroller)
	if err != nil {
		return nil, err
	}
	return Microcontroller, nil
}

func (re *MicrocontrollerRepository) Get(id int64) (*entity.Microcontroller, error) {

	Microcontroller := service.FactoryMicrocontroller()
	re.First(Microcontroller, id)
	return Microcontroller, nil
}

func (re *MicrocontrollerRepository) GetAll(conditions ...interface{}) []*entity.Microcontroller {

	var Microcontrollers []*entity.Microcontroller
	re.Find(&Microcontrollers, conditions...)
	return Microcontrollers
}

func (re *MicrocontrollerRepository) Save(Microcontroller *entity.Microcontroller) (*entity.Microcontroller, error) {

	if Microcontroller.Id == 0 {
		_, err := re.create(Microcontroller)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := re.update(Microcontroller)
		if err != nil {
			return nil, err
		}
	}
	return Microcontroller, nil
}

func (re *MicrocontrollerRepository) Remove(Microcontroller *entity.Microcontroller) error {

	result := re.Delete(Microcontroller)
	if result != nil {
		return result
	}
	return nil
}

func (re *MicrocontrollerRepository) SqlQuery(sql string, values ...interface{}) error {
	return re.Exec(sql, values...)
}

func (re *MicrocontrollerRepository) SqlQueryRow(sql string) *sql.Row {
	return re.QueryRow(sql)
}

func (re *MicrocontrollerRepository) SqlQueryRows(sql string) (*sql.Rows, error) {
	return re.QueryRows(sql)
}

func (re *MicrocontrollerRepository) SqlQueryPaginator(columns string, table string, where string, sqlTemplate string, page float64, limit float64, order ...string) (map[string]interface{}, error) {
	return re.QueryPaginator(columns, table, where, sqlTemplate, page, limit, order...)
}
