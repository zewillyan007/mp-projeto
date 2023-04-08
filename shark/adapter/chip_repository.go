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

type ChipRepository struct {
	*adapter.Repository
}

func NewChipRepository(db *gorm.DB) *ChipRepository {

	return &ChipRepository{
		Repository: &adapter.Repository{Db: db, Table: "shark.chip"},
	}
}

func (re *ChipRepository) WithTransaction(transaction port_shared.ITransaction) port.ChipIRepository {
	return NewChipRepository(transaction.GetTransaction())
}

func (re *ChipRepository) create(Chip *entity.Chip) (interface{}, error) {

	result := re.Insert(Chip)
	if result != nil {
		return nil, result
	}
	return Chip, nil
}

func (re *ChipRepository) update(Chip *entity.Chip) (interface{}, error) {

	err := re.Update(Chip)
	if err != nil {
		return nil, err
	}
	return Chip, nil
}

func (re *ChipRepository) Get(id int64) (*entity.Chip, error) {

	Chip := service.FactoryChip()
	re.First(Chip, id)
	return Chip, nil
}

func (re *ChipRepository) GetAll(conditions ...interface{}) []*entity.Chip {

	var Chips []*entity.Chip
	re.Find(&Chips, conditions...)
	return Chips
}

func (re *ChipRepository) Save(Chip *entity.Chip) (*entity.Chip, error) {

	if Chip.Id == 0 {
		_, err := re.create(Chip)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := re.update(Chip)
		if err != nil {
			return nil, err
		}
	}
	return Chip, nil
}

func (re *ChipRepository) Remove(Chip *entity.Chip) error {

	result := re.Delete(Chip)
	if result != nil {
		return result
	}
	return nil
}

func (re *ChipRepository) SqlQuery(sql string, values ...interface{}) error {
	return re.Exec(sql, values...)
}

func (re *ChipRepository) SqlQueryRow(sql string) *sql.Row {
	return re.QueryRow(sql)
}

func (re *ChipRepository) SqlQueryRows(sql string) (*sql.Rows, error) {
	return re.QueryRows(sql)
}

func (re *ChipRepository) SqlQueryPaginator(columns string, table string, where string, sqlTemplate string, page float64, limit float64, order ...string) (map[string]interface{}, error) {
	return re.QueryPaginator(columns, table, where, sqlTemplate, page, limit, order...)
}
