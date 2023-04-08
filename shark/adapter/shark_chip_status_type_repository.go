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

type SharkChipStatusTypeRepository struct {
	*adapter.Repository
}

func NewSharkChipStatusTypeRepository(db *gorm.DB) *SharkChipStatusTypeRepository {

	return &SharkChipStatusTypeRepository{
		Repository: &adapter.Repository{Db: db, Table: "domain.shark_chip_status_type"},
	}
}

func (re *SharkChipStatusTypeRepository) WithTransaction(transaction port_shared.ITransaction) port.SharkChipStatusTypeIRepository {
	return NewSharkChipStatusTypeRepository(transaction.GetTransaction())
}

func (re *SharkChipStatusTypeRepository) create(SharkChipStatusType *entity.SharkChipStatusType) (interface{}, error) {

	result := re.Insert(SharkChipStatusType)
	if result != nil {
		return nil, result
	}
	return SharkChipStatusType, nil
}

func (re *SharkChipStatusTypeRepository) update(SharkChipStatusType *entity.SharkChipStatusType) (interface{}, error) {

	err := re.Update(SharkChipStatusType)
	if err != nil {
		return nil, err
	}
	return SharkChipStatusType, nil
}

func (re *SharkChipStatusTypeRepository) Get(id int64) (*entity.SharkChipStatusType, error) {

	SharkChipStatusType := service.FactorySharkChipStatusType()
	re.First(SharkChipStatusType, id)
	return SharkChipStatusType, nil
}

func (re *SharkChipStatusTypeRepository) GetAll(conditions ...interface{}) []*entity.SharkChipStatusType {

	var SharkChipStatusTypes []*entity.SharkChipStatusType
	re.Find(&SharkChipStatusTypes, conditions...)
	return SharkChipStatusTypes
}

func (re *SharkChipStatusTypeRepository) Save(SharkChipStatusType *entity.SharkChipStatusType) (*entity.SharkChipStatusType, error) {

	if SharkChipStatusType.Id == 0 {
		_, err := re.create(SharkChipStatusType)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := re.update(SharkChipStatusType)
		if err != nil {
			return nil, err
		}
	}
	return SharkChipStatusType, nil
}

func (re *SharkChipStatusTypeRepository) Remove(SharkChipStatusType *entity.SharkChipStatusType) error {

	result := re.Delete(SharkChipStatusType)
	if result != nil {
		return result
	}
	return nil
}

func (re *SharkChipStatusTypeRepository) SqlQuery(sql string, values ...interface{}) error {
	return re.Exec(sql, values...)
}

func (re *SharkChipStatusTypeRepository) SqlQueryRow(sql string) *sql.Row {
	return re.QueryRow(sql)
}

func (re *SharkChipStatusTypeRepository) SqlQueryRows(sql string) (*sql.Rows, error) {
	return re.QueryRows(sql)
}

func (re *SharkChipStatusTypeRepository) SqlQueryPaginator(columns string, table string, where string, sqlTemplate string, page float64, limit float64, order ...string) (map[string]interface{}, error) {
	return re.QueryPaginator(columns, table, where, sqlTemplate, page, limit, order...)
}
