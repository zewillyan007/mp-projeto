package grid

type GridParam struct {
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

type GridOrder map[string]string

type GridConfig struct {
	Page     string                   `json:"page"`
	RowsPage string                   `json:"rows_page"`
	Params   []map[string][]GridParam `json:"params"`
	Orderby  []GridOrder              `json:"orderby"`
}

func NewGridConfig() *GridConfig {
	return &GridConfig{
		Params:  []map[string][]GridParam{},
		Orderby: []GridOrder{},
	}
}
