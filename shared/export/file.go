package export

import (
	"bytes"
	"encoding/csv"
)

const (
	TYPE_CSV = "csv"
	TYPE_PDF = "pdf"
)

func FileExport(fileType string, data [][]string) ([]byte, error) {

	var err error
	var bytes []byte

	switch fileType {
	case TYPE_CSV:
		bytes, err = csvFile(data)
	case TYPE_PDF:
		bytes, err = pdfFile(data)
	}

	return bytes, err
}

func csvFile(data [][]string) ([]byte, error) {

	var b []byte
	buffer := bytes.NewBuffer(b)
	w := csv.NewWriter(buffer)
	for _, row := range data {
		if err := w.Write(row); err != nil {
			return nil, err
		}
	}
	return buffer.Bytes(), w.Error()
}

func pdfFile(data [][]string) ([]byte, error) {

	return nil, nil
}
