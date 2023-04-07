package adapter

import (
	"database/sql"
	"fmt"
	"math"
	"mp-projeto/shared/grid"
	"reflect"
	"strings"

	"gorm.io/gorm"
)

type Repository struct {
	Db    *gorm.DB
	Table string
}

func NewRepository(db *gorm.DB, table string) *Repository {
	return &Repository{Db: db, Table: table}
}

func (o *Repository) GetTable() string {
	return o.Table
}

func (o *Repository) Insert(value interface{}) error {
	result := o.Db.Table(o.Table).Create(value)
	return result.Error
}

func (o *Repository) Update(value interface{}) error {
	result := o.Db.Table(o.Table).Save(value)
	return result.Error
}

func (o *Repository) Delete(value interface{}) error {
	result := o.Db.Table(o.Table).Delete(value)
	return result.Error
}

func (o *Repository) First(dest interface{}, conds ...interface{}) {
	if conds != nil {
		o.Db.Table(o.Table).Find(dest, conds)
	} else {
		o.Db.Table(o.Table).Find(dest)
	}
}

func (o *Repository) Find(dest interface{}, conds ...interface{}) {
	o.Db.Table(o.Table).Find(dest, conds...)
}

func (o *Repository) Exec(sql string, values ...interface{}) error {
	result := o.Db.Exec(sql, values...)
	return result.Error
}

func (o *Repository) QueryRow(sql string) *sql.Row {
	return o.Db.Raw(sql).Row()
}

func (o *Repository) QueryRows(sql string) (*sql.Rows, error) {
	return o.Db.Raw(sql).Rows()
}

func (o *Repository) QueryPaginatorGroupBy(columns, table, where, group, sqlTemplate string, page float64, limit float64, order ...string) (map[string]interface{}, error) {
	if len(group) > 0 {
		sqlTemplate = sqlTemplate + " GROUP BY " + group
	}
	return o.QueryPaginator(columns, table, where, sqlTemplate, page, limit, order...)
}

func (o *Repository) QueryPaginator(columns, table, where, sqlTemplate string, page float64, limit float64, order ...string) (map[string]interface{}, error) {

	var sql string
	var orderby string
	var total, offset, pages float64
	data := make(map[string]interface{})

	var sqlParsed string
	if len(where) > 0 {
		sqlParsed = fmt.Sprintf(sqlTemplate, columns, table, where)
	} else {
		sqlParsed = fmt.Sprintf(sqlTemplate, columns, table)
	}

	if limit > 0 {
		offset = (page - 1) * limit
		rowsCount := o.QueryRow("SELECT COUNT(*) FROM " + table + " WHERE " + where)
		rowsCount.Scan(&total)
		if len(order) > 0 && len(order[0]) > 0 {
			orderby = " ORDER BY " + order[0]
		}
		sql = sqlParsed + orderby + " OFFSET " + fmt.Sprintf("%v", offset) + " LIMIT " + fmt.Sprintf("%v", limit)
	} else {
		sql = sqlParsed + orderby
	}

	rows, err := o.QueryRows(sql)

	if err != nil {
		return nil, err
	}

	if limit > 0 {
		pages = math.Ceil(total / limit)
		if page > 1 {
			data["prev"] = page - 1
		} else {
			data["prev"] = page
		}
		if float64(page+1) <= pages {
			data["next"] = page + 1
		} else {
			data["next"] = page
		}
	} else {
		pages = 0
		data["prev"] = 0
		data["next"] = 0
	}

	data["page"] = page
	data["pages"] = pages
	data["total"] = total
	data["rows"] = grid.RowsToInterface(rows)
	data["lines"] = len(data["rows"].([]interface{}))

	return data, nil
}

func (o *Repository) WhereCondition(condition interface{}) string {

	var expr string = ""
	cond := condition.([]interface{})

	if len(cond) > 0 {
		expr = cond[0].(string)
		values := cond[1:]
		for _, val := range values {
			expr = strings.Replace(expr, "?", "%v", 1)
			switch reflect.TypeOf(val).Kind().String() {
			case "string":
				expr = fmt.Sprintf(expr, val.(string))
			case "int":
				expr = fmt.Sprintf(expr, val.(int))
			case "int32":
				expr = fmt.Sprintf(expr, val.(int))
			case "int64":
				expr = fmt.Sprintf(expr, val.(int64))
			case "float32":
				expr = fmt.Sprintf(expr, val.(float32))
			case "float64":
				expr = fmt.Sprintf(expr, val.(float64))
			}
		}
	}

	return expr
}
