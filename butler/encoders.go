package butler

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	html "html/template"
	"reflect"
	"strings"
	text "text/template"
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
		tagNames := tag.Get(DefaultCsvEncoderNamespace)

		if tagNames != "" {
			parts := strings.Split(tagNames, ",")
			tagName := parts[0]

			value := fmt.Sprintf("%v", valueField.Interface())
			if len(parts) > 0 && parts[1] == "omitempty" && value == "" {
				continue
			}

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

type TextEncoder struct {
	Template string
}

func (e TextEncoder) Encode(a Any) ([]byte, error) {
	var (
		buffer *bytes.Buffer
	)
	tmpl := text.Must(text.New("text-encoder").Parse(e.Template))
	if err := tmpl.Execute(buffer, a); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

type HtmlEncoder struct {
	Template string
}

func (e HtmlEncoder) Encode(a Any) ([]byte, error) {
	var (
		buffer *bytes.Buffer
	)
	tmpl, err := html.New("html-encoder").Parse(e.Template)
	if err != nil {
		return nil, err
	}
	if err := tmpl.Execute(buffer, a); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

type XmlEncoder struct{}

func (e XmlEncoder) Encode(a Any) ([]byte, error) {
	return xml.Marshal(a)
}
