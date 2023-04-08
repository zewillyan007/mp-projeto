package usecase

import (
	"mp-projeto/shared/grid"
	"mp-projeto/shark/core/port"
	"strconv"
)

type IncidenceUseCaseGrid struct {
	repository port.IncidenceIRepository
}

func NewIncidenceUseCaseGrid(repository port.IncidenceIRepository) *IncidenceUseCaseGrid {
	return &IncidenceUseCaseGrid{repository: repository}
}

func (o *IncidenceUseCaseGrid) Execute(GridConfig *grid.GridConfig) map[string]interface{} {

	var page, limit float64
	var columns, where, order, table, sql string

	params := grid.NewParams()
	orders := grid.NewOrders()

	params.LoadGridParams(GridConfig)
	orders.LoadGridOrders(GridConfig)

	columns = "*"
	table = "shark.incidence"
	order = ""
	where = params.ToString()

	if len(where) == 0 {
		where = "1 = 1"
	}

	if len(orders.ToString()) > 0 {
		order = orders.ToString()
	}

	sql = "SELECT %s FROM %s WHERE %s"
	page, _ = strconv.ParseFloat(GridConfig.Page, 64)
	limit, _ = strconv.ParseFloat(GridConfig.RowsPage, 64)
	data, err := o.repository.SqlQueryPaginator(columns, table, where, sql, page, limit, order)

	if err != nil {
		panic(err.Error())
	}

	return data
}
