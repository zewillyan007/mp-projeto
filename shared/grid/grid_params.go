package grid

import (
	"fmt"
	"reflect"
	"strings"
)

type Params struct {
	search      map[string]string
	list        map[string][]map[string]interface{}
	list_search map[string][]map[string]interface{}
}

func NewParams() *Params {
	return &Params{
		list:        map[string][]map[string]interface{}{},
		list_search: map[string][]map[string]interface{}{},
	}
}

func (o *Params) Validate(columns []string) []string {

	var invalidFields []string
	listFilter := o.GetList()
	filterKeys := o.GetListKeys()
	columnsString := strings.Join(columns, ",")

	for _, filter := range filterKeys {
		if !strings.Contains(columnsString, filter) {
			invalidFields = append(invalidFields, filter)
		}

		// reflect.ValueOf(a).IsNil()
		for _, field := range listFilter[filter] {

			if reflect.TypeOf(field["value"]).Kind().String() == "string" && field["value"].(string) == "" {
				invalidFields = append(invalidFields, filter)
			} else {
				if field["value"] == nil {
					invalidFields = append(invalidFields, filter)
				}
			}
		}
	}
	return invalidFields
}

func (o *Params) ValidateMandatory(mandatory []string) []string {

	var emptyFields []string
	filterKeysString := strings.Join(o.GetListKeys(), ",")

	for _, filter := range mandatory {
		if !strings.Contains(filterKeysString, filter) {
			emptyFields = append(emptyFields, filter)
		}
	}
	return emptyFields
}

func (o *Params) GetList() map[string][]map[string]interface{} {
	return o.list
}

func (o *Params) GetListKeys() []string {

	var listKeys []string
	for key := range o.list {
		listKeys = append(listKeys, key)
	}
	return listKeys
}

func (o *Params) GetListSearch() map[string][]map[string]interface{} {
	return o.list_search
}

func (o *Params) SetSearchFields(fields map[string]string) {
	o.search = fields
}

func (o *Params) LoadGridParams(grid *GridConfig) {

	for _, v := range grid.Params {
		for field, v2 := range v {
			if field != "_search_" {
				for _, v3 := range v2 {
					o.Add(field, v3.Operator, v3.Value)
				}
			} else {
				if o.search != nil && len(o.search) > 0 {
					for column := range o.search {
						o.AddSearch(column, v2[0].Operator, v2[0].Value)
					}
				}
			}
		}
	}
}

func (o *Params) Add(field string, operator string, value interface{}) {

	param := make(map[string]interface{})
	param["operator"] = operator
	param["value"] = value
	o.list[field] = append(o.list[field], param)
}

func (o *Params) AddSearch(field string, operator string, value interface{}) {

	param := make(map[string]interface{})
	param["operator"] = operator
	param["value"] = value
	o.list_search[field] = append(o.list_search[field], param)
}

func (o *Params) ToArrayParams() []string {

	var where []string
	parts := make(map[string][]string)
	for field, sublist := range o.list {
		for _, values := range sublist {
			switch reflect.TypeOf(values["value"]).Kind().String() {
			case "int", "int32", "int64", "float", "float32", "float64":
				parts[field] = append(parts[field], field+" "+values["operator"].(string)+" "+fmt.Sprintf("%v", values["value"]))
			default:
				parts[field] = append(parts[field], field+" "+values["operator"].(string)+" '"+fmt.Sprintf("%v", values["value"])+"'")
			}
		}
	}
	for _, v := range parts {
		where = append(where, "("+strings.Join(v, " OR ")+")")
	}
	return where
}

func (o *Params) ToArraySearch() []string {

	var where []string
	parts := make(map[string][]string)
	for field, sublist := range o.list_search {
		for _, values := range sublist {
			switch o.search[field] {
			case "numeric":
				switch reflect.TypeOf(values["value"]).Kind().String() {
				case "int", "int32", "int64", "float", "float32", "float64":
					parts[field] = append(parts[field], field+" = "+fmt.Sprintf("%v", values["value"]))
				}
			case "string":
				if strings.ToUpper(values["operator"].(string)) == "ILIKE" || strings.ToUpper(values["operator"].(string)) == "LIKE" {
					parts[field] = append(parts[field], field+" "+values["operator"].(string)+" '%"+fmt.Sprintf("%v", values["value"])+"%'")
				} else {
					parts[field] = append(parts[field], field+" "+values["operator"].(string)+" '"+fmt.Sprintf("%v", values["value"])+"'")
				}
			}
		}
	}
	for _, v := range parts {
		where = append(where, "("+strings.Join(v, " OR ")+")")
	}
	return where
}

func (o *Params) ToString() string {
	return strings.Join(o.ToArrayParams(), " AND ")
}

func (o *Params) ToStringSearch() string {
	return strings.Join(o.ToArraySearch(), " OR ")
}
