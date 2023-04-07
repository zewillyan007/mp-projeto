package grid

import (
	"strings"
)

type Orders struct {
	list map[string]string
}

func NewOrders() *Orders {
	return &Orders{
		list: map[string]string{},
	}
}

func (o *Orders) Validate(columns []string) []string {

	var invalidFields []string
	filterKeys := o.GetListKeys()
	columnsString := strings.Join(columns, ",")

	for _, filter := range filterKeys {
		if !strings.Contains(columnsString, filter) {
			invalidFields = append(invalidFields, filter)
		}
	}
	return invalidFields
}

func (o *Orders) GetList() map[string]string {
	return o.list
}

func (o *Orders) GetListKeys() []string {

	var listKeys []string
	for key := range o.list {
		listKeys = append(listKeys, key)
	}
	return listKeys
}

func (o *Orders) LoadGridOrders(grid *GridConfig) {
	for _, v := range grid.Orderby {
		for field, value := range v {
			o.Add(strings.ToLower(field), value)
		}
	}
}

func (o *Orders) ToString() string {
	return strings.Join(o.ToArrayParams(), ", ")
}

func (o *Orders) ToStringTranslate(dictionaryType map[string]string) string {
	return strings.Join(o.ToArrayParamsTranslate(dictionaryType), ",")
}

func (o *Orders) Add(field string, direction string) {
	o.list[field] = direction
}

func (o *Orders) ToArrayParams() []string {

	var order []string
	for field, direction := range o.list {
		order = append(order, field+" "+direction)
	}
	return order
}

func (o *Orders) ToArrayParamsTranslate(dictionaryType map[string]string) []string {

	var order []string
	for field, direction := range o.list {
		if dictionaryType[field] == "string" {
			order = append(order, "translate(lower("+field+"),'áéíóúýçôêõãüöëñ','aeiouycoeoauoen') "+direction)
		} else {
			order = append(order, field+" "+direction)
		}
	}
	return order
}
