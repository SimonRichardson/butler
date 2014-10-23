package butler

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"reflect"
)

type Encoder interface {
	Encode(a Any) []byte
}

const (
	DefaultCsvEncoderNamespace string = "csv"
)

type CsvEncoder struct{}

func (e CsvEncoder) Encode(a Any) ([]byte, error) {
	var (
		buffer  *bytes.Buffer
		headers []string
		values  []string
	)
	writer := csv.NewWriter(buffer)

	val := reflect.ValueOf(a).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		tag := typeField.Tag
		tagName := tag.Get(DefaultCsvEncoderNamespace)

		if tagName != "" {
			value := fmt.Sprintf("%v", valueField.Interface())

			headers = append(headers, tagName)
			values = append(values, value)
		}
	}

	if err := writer.Write(headers); err != nil {
		return nil, err
	}
	if err := writer.Write(values); err != nil {
		return nil, err
	}

	writer.Flush()
	return buffer.Bytes(), nil
}

type JsonEncoder struct{}

func (e JsonEncoder) Encode(a Any) ([]byte, error) {
	return json.Marshal(a)
}

type XmlEncoder struct{}

func (e XmlEncoder) Encode(a Any) ([]byte, error) {
	return xml.Marshal(a)
}
