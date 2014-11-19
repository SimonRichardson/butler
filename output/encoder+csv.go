package output

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"reflect"
	"strings"

	g "github.com/SimonRichardson/butler/generic"
)

const (
	DefaultCsvEncoderNamespace string = "csv"
)

type CsvEncoder struct{}

func (e CsvEncoder) Encode(a g.Any) g.Either {
	// This is horrid, re-work it!
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
		return g.NewLeft(err)
	}
	if err := writer.Write(values); err != nil {
		return g.NewLeft(err)
	}

	writer.Flush()
	return g.NewRight(buffer.Bytes())
}

func (e CsvEncoder) Generate(x g.Any) g.Either {
	return generate(e)(x)
}
