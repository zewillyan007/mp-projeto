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

type SharkChipRepository struct {
	*adapter.Repository
}

func NewSharkChipRepository(db *gorm.DB) *SharkChipRepository {

	return &SharkChipRepository{
		Repository: &adapter.Repository{Db: db, Table: "shark.shark_chip"},
	}
}

func (re *SharkChipRepository) WithTransaction(transaction port_shared.ITransaction) port.SharkChipIRepository {
	return NewSharkChipRepository(transaction.GetTransaction())
}

func (re *SharkChipRepository) create(SharkChip *entity.SharkChip) (interface{}, error) {

	result := re.Insert(SharkChip)
	if result != nil {
		return nil, result
	}
	return SharkChip, nil
}

func (re *SharkChipRepository) update(SharkChip *entity.SharkChip) (interface{}, error) {

	err := re.Update(SharkChip)
	if err != nil {
		return nil, err
	}
	return SharkChip, nil
}

func (re *SharkChipRepository) Get(id int64) (*entity.SharkChip, error) {

	SharkChip := service.FactorySharkChip()
	re.First(SharkChip, id)
	return SharkChip, nil
}

func (re *SharkChipRepository) GetAll(conditions ...interface{}) []*entity.SharkChip {

	var SharkChips []*entity.SharkChip
	re.Find(&SharkChips, conditions...)
	return SharkChips
}

func (re *SharkChipRepository) Save(SharkChip *entity.SharkChip) (*entity.SharkChip, error) {

	if SharkChip.Id == 0 {
		_, err := re.create(SharkChip)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := re.update(SharkChip)
		if err != nil {
			return nil, err
		}
	}
	return SharkChip, nil
}

func (re *SharkChipRepository) Remove(SharkChip *entity.SharkChip) error {

	result := re.Delete(SharkChip)
	if result != nil {
		return result
	}
	return nil
}

func (re *SharkChipRepository) SqlQuery(sql string, values ...interface{}) error {
	return re.Exec(sql, values...)
}

func (re *SharkChipRepository) SqlQueryRow(sql string) *sql.Row {
	return re.QueryRow(sql)
}

func (re *SharkChipRepository) SqlQueryRows(sql string) (*sql.Rows, error) {
	return re.QueryRows(sql)
}

func (re *SharkChipRepository) SqlQueryPaginator(columns string, table string, where string, sqlTemplate string, page float64, limit float64, order ...string) (map[string]interface{}, error) {
	return re.QueryPaginator(columns, table, where, sqlTemplate, page, limit, order...)
}
