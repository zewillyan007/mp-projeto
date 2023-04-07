package grid

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Grid struct {
}

func (o *Grid) Prepare(GridConfig *GridConfig, columns, mandatory []string, search map[string]string) (map[string]interface{}, error) {

	prep := make(map[string]interface{}, 0)
	params := NewParams()
	orders := NewOrders()
	params.SetSearchFields(search)
	params.LoadGridParams(GridConfig)
	orders.LoadGridOrders(GridConfig)
	invalidFields := make(map[string][]string, 0)

	mandatoryParams := params.ValidateMandatory(mandatory)
	if len(mandatoryParams) > 0 {
		invalidFields["mandatory"] = mandatoryParams
	}

	invalidParams := params.Validate(columns)
	if len(invalidParams) > 0 {
		invalidFields["params"] = invalidParams
	}

	invalidOrders := orders.Validate(columns)
	if len(invalidOrders) > 0 {
		invalidFields["orders"] = invalidOrders
	}

	if len(invalidFields) > 0 {
		msg := make(map[string]string, 0)
		if len(invalidFields["mandatory"]) > 0 {
			msg["mandatory"] = strings.Join(invalidFields["mandatory"], ", ")
		}
		if len(invalidFields["params"]) > 0 {
			msg["params"] = strings.Join(invalidFields["params"], ", ")
		}
		if len(invalidFields["orders"]) > 0 {
			msg["orders"] = strings.Join(invalidFields["orders"], ", ")
		}
		_json_, _ := json.Marshal(msg)
		return nil, errors.New(string(_json_))
	}

	prep["params"] = params
	prep["orders"] = orders

	return prep, nil
}

func (o *Grid) ParamIntervalDate(params *Params, nameInitDate, nameFinDate string) map[string]interface{} {

	mapa := make(map[string]interface{})
	paramsNew := NewParams()
	listParams := params.GetList()
	for nameParam, valueParam := range listParams {
		if nameParam != nameInitDate && nameParam != nameFinDate {
			for _, val := range valueParam {
				paramsNew.Add(nameParam, val["operator"].(string), val["value"])
			}
		} else {
			if nameParam == nameInitDate {
				mapa[nameInitDate] = valueParam[0]["value"].(string)
			}
			if nameParam == nameFinDate {
				mapa[nameFinDate] = valueParam[0]["value"].(string)
			}
		}
	}
	mapa["params"] = paramsNew
	return mapa
}

func (o *Grid) VisionCompaniesAndRegionalsArrayToString(params *Params, companies, regionals []int64) (string, string) {

	var strCompanies string
	if len(companies) > 0 {
		strCompanies = fmt.Sprintf("(id_company in (%s))", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(companies)), ","), "[]"))
	}

	var strRegionals string
	if len(regionals) > 0 {
		strRegionals = fmt.Sprintf("(id_regional in (%s))", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(regionals)), ","), "[]"))
	}

	return strCompanies, strRegionals
}

func (o *Grid) ArrayToString(params *Params, fieldName string, array []int64) string {

	paramsList := params.GetList()

	var out string
	if len(array) > 0 {
		if _, found := paramsList[fieldName]; !found {
			out = fmt.Sprintf("("+fieldName+"in (%s))", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(array)), ","), "[]"))
		}
	}

	return out
}
