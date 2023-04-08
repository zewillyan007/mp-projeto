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

type ChipStatusTypeRepository struct {
	*adapter.Repository
}

func NewChipStatusTypeRepository(db *gorm.DB) *ChipStatusTypeRepository {

	return &ChipStatusTypeRepository{
		Repository: &adapter.Repository{Db: db, Table: "domain.chip_status_type"},
	}
}

func (re *ChipStatusTypeRepository) WithTransaction(transaction port_shared.ITransaction) port.ChipStatusTypeIRepository {
	return NewChipStatusTypeRepository(transaction.GetTransaction())
}

func (re *ChipStatusTypeRepository) create(ChipStatusType *entity.ChipStatusType) (interface{}, error) {

	result := re.Insert(ChipStatusType)
	if result != nil {
		return nil, result
	}
	return ChipStatusType, nil
}

func (re *ChipStatusTypeRepository) update(ChipStatusType *entity.ChipStatusType) (interface{}, error) {

	err := re.Update(ChipStatusType)
	if err != nil {
		return nil, err
	}
	return ChipStatusType, nil
}

func (re *ChipStatusTypeRepository) Get(id int64) (*entity.ChipStatusType, error) {

	ChipStatusType := service.FactoryChipStatusType()
	re.First(ChipStatusType, id)
	return ChipStatusType, nil
}

func (re *ChipStatusTypeRepository) GetAll(conditions ...interface{}) []*entity.ChipStatusType {

	var ChipStatusTypes []*entity.ChipStatusType
	re.Find(&ChipStatusTypes, conditions...)
	return ChipStatusTypes
}

func (re *ChipStatusTypeRepository) Save(ChipStatusType *entity.ChipStatusType) (*entity.ChipStatusType, error) {

	if ChipStatusType.Id == 0 {
		_, err := re.create(ChipStatusType)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := re.update(ChipStatusType)
		if err != nil {
			return nil, err
		}
	}
	return ChipStatusType, nil
}

func (re *ChipStatusTypeRepository) Remove(ChipStatusType *entity.ChipStatusType) error {

	result := re.Delete(ChipStatusType)
	if result != nil {
		return result
	}
	return nil
}

func (re *ChipStatusTypeRepository) SqlQuery(sql string, values ...interface{}) error {
	return re.Exec(sql, values...)
}

func (re *ChipStatusTypeRepository) SqlQueryRow(sql string) *sql.Row {
	return re.QueryRow(sql)
}

func (re *ChipStatusTypeRepository) SqlQueryRows(sql string) (*sql.Rows, error) {
	return re.QueryRows(sql)
}

func (re *ChipStatusTypeRepository) SqlQueryPaginator(columns string, table string, where string, sqlTemplate string, page float64, limit float64, order ...string) (map[string]interface{}, error) {
	return re.QueryPaginator(columns, table, where, sqlTemplate, page, limit, order...)
}
