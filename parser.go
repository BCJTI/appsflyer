package appsflyer

import (
	"encoding/csv"
	"io"
	"reflect"
	"strings"
)

func getReader(body []byte) *csv.Reader {
	reader := csv.NewReader(strings.NewReader(strings.TrimSuffix(string(body), "\n")))
	reader.FieldsPerRecord = -1
	return reader
}

func Parse(body []byte, v interface{}, f func(result interface{})) error {

	reader := getReader(body)
	headers, err := reader.Read()
	if err != nil {
		return err
	}

	t := reflect.TypeOf(v)
	newData := reflect.New(t).Elem()

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		for i := 0; i < t.NumField(); i++ {
			var (
				f     = t.Field(i)
				index = -1
			)
			fieldName := f.Tag.Get("csv")
			for k, v := range headers {
				if v == fieldName {
					index = k
				}
			}
			if index == -1 {
				continue
			}
			newField := newData.FieldByName(f.Name)
			if !newField.IsValid() || !newField.CanSet() {
				continue
			}
			newField.Set(reflect.ValueOf(row[index]))
		}
		f(newData.Interface())
	}
	return nil
}

func Map(body []byte) (rows []map[string]interface{}, err error) {

	var header []string

	// get csv reader
	reader := getReader(body)

	// get header (first row)
	if header, err = reader.Read(); err != nil {
		return
	}

	for {

		var line []string

		// get csv line
		line, err = reader.Read()

		if err != nil {

			// reset error
			if err == io.EOF {
				err = nil
			}

			break

		}

		// init report record
		record := make(map[string]interface{})

		// parse data to correct datatype
		for i, column := range header {
			record[column] = StringToInterface(line[i])
		}

		// append record
		rows = append(rows,record)

	}

	return

}