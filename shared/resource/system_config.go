package resource

import (
	"database/sql"
	"strconv"
)

type Config struct {
	Id            int64  `json:"id"`
	Key           string `json:"key"`
	Value         string `json:"value"`
	Description   string `json:"description"`
	Active        string `json:"active"`
	IdSaleNetwork int64  `json:"id_sale_network"`
}

type SystemConfig struct {
	configs map[string]*Config
}

func NewConfig() *Config {
	return &Config{}
}

func NewSystemConfig() *SystemConfig {
	return &SystemConfig{
		configs: map[string]*Config{},
	}
}

func (o *SystemConfig) Load(rows *sql.Rows) {

	for _, c := range o.RowsToInterface(rows) {
		m := c.(map[string]interface{})
		id, _ := strconv.ParseInt(m["id"].(string), 10, 0)
		idSaleNetwork, _ := strconv.ParseInt(m["id_sale_network"].(string), 10, 0)
		config := NewConfig()
		config.Id = id
		config.Key = m["key"].(string)
		config.Value = m["value"].(string)
		config.Active = m["active"].(string)
		config.Description = m["description"].(string)
		config.IdSaleNetwork = idSaleNetwork
		o.configs[config.Key] = config
	}
}

func (o *SystemConfig) Get(key string) *Config {

	return o.configs[key]
}

func (*SystemConfig) RowsToInterface(rows *sql.Rows) []interface{} {

	defer rows.Close()
	columnTypes, _ := rows.ColumnTypes()
	count := len(columnTypes)
	finalRows := []interface{}{}

	for rows.Next() {
		scanArgs := make([]interface{}, count)
		for i, v := range columnTypes {
			switch v.DatabaseTypeName() {
			case "VARCHAR", "TEXT", "UUID", "TIMESTAMP":
				scanArgs[i] = new(sql.NullString)
			case "BOOL":
				scanArgs[i] = new(sql.NullBool)
			case "INT4":
				scanArgs[i] = new(sql.NullInt64)
			default:
				scanArgs[i] = new(sql.NullString)
			}
		}

		rows.Scan(scanArgs...)
		masterData := map[string]interface{}{}

		for i, v := range columnTypes {
			if z, ok := (scanArgs[i]).(*sql.NullBool); ok {
				masterData[v.Name()] = z.Bool
				continue
			}
			if z, ok := (scanArgs[i]).(*sql.NullString); ok {
				masterData[v.Name()] = z.String
				continue
			}
			if z, ok := (scanArgs[i]).(*sql.NullInt64); ok {
				masterData[v.Name()] = z.Int64
				continue
			}
			if z, ok := (scanArgs[i]).(*sql.NullFloat64); ok {
				masterData[v.Name()] = z.Float64
				continue
			}
			if z, ok := (scanArgs[i]).(*sql.NullInt32); ok {
				masterData[v.Name()] = z.Int32
				continue
			}
			masterData[v.Name()] = scanArgs[i]
		}
		finalRows = append(finalRows, masterData)
	}

	return finalRows
}
