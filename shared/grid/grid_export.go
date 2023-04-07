package grid

import (
	"mp-projeto/shared/export"
	"net/http"
	"strconv"
	"strings"
)

var CONTENT_TYPE_DATA map[string]string = map[string]string{
	"csv": "text/csv",
}

func ExportDataGrid(fileType string, grid map[string]interface{}) ([]byte, error) {

	var row []string
	var data [][]string
	var columns []string

	for _, dataRow := range grid["rows"].([]interface{}) {
		row = make([]string, 0)
		cols := make([]string, 0)
		for column, value := range dataRow.(map[string]interface{}) {
			cols = append(cols, strings.ToUpper(column))
			row = append(row, value.(string))
		}
		if len(columns) == 0 {
			columns = cols
			data = append(data, columns)
		}
		data = append(data, row)
	}

	return export.FileExport(fileType, data)
}

func SetHeaderType(w http.ResponseWriter, fileType string, bytes []byte, fileName ...string) {

	if len(fileName) > 0 {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Length", strconv.Itoa(len(bytes)))
		w.Header().Set("Content-Disposition", "attachment; filename="+fileName[0]+"."+strings.ToLower(fileType))
	} else {
		w.Header().Set("Content-Type", CONTENT_TYPE_DATA[strings.ToLower(fileType)])
		w.Header().Set("Content-Length", strconv.Itoa(len(bytes)))
	}
}

func ResponseDataGrid(w http.ResponseWriter, fileType string, grid map[string]interface{}, fileName ...string) error {

	bytes, err := ExportDataGrid(fileType, grid)
	SetHeaderType(w, fileType, bytes, fileName...)
	_, err = w.Write(bytes)
	return err
}
