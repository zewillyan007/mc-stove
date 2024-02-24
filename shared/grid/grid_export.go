package grid

import (
	"mc-stove/shared/export"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

var CONTENT_TYPE_DATA map[string]string = map[string]string{
	"csv": "text/csv",
}

func ExportDataGrid(fileType string, grid map[string]interface{}) ([]byte, error) {

	var row []string
	var data [][]string

	//GET COLUMNS
	firstRow := grid["rows"].([]interface{})[0].(map[string]interface{})
	keys := make([]string, 0, len(firstRow))
	for k := range firstRow {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	data = append(data, keys)

	//INSERT VALUES
	for _, dataRow := range grid["rows"].([]interface{}) {
		row = make([]string, 0)
		for _, k2 := range keys {
			r1 := dataRow.(map[string]interface{})
			row = append(row, r1[k2].(string))
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
	if err != nil {
		return err
	}

	SetHeaderType(w, fileType, bytes, fileName...)
	_, err = w.Write(bytes)
	return err
}
