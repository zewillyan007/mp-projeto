package grid

type GridData struct {
	Lines int
	Next  int
	Page  int
	Pages int
	Prev  int
	Rows  []interface{}
	Total int
}

func NewGridData() *GridData {
	return &GridData{
		Lines: 0,
		Next:  1,
		Page:  1,
		Pages: 1,
		Prev:  1,
		Rows:  []interface{}{},
		Total: 0,
	}
}
