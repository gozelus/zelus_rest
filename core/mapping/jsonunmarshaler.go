package mapping

import (
	"github.com/gozelus/zelus_rest/core/jsonz"
	"io"
)

const jsonTagKey = "json"

var jsonUnmarshaler = NewUnmarshaler(jsonTagKey)

func UnmarshalJsonBytes(content []byte, v interface{}) error {
	return unmarshalJsonBytes(content, v, jsonUnmarshaler)
}

func UnmarshalJsonReader(reader io.Reader, v interface{}) error {
	return unmarshalJsonReader(reader, v, jsonUnmarshaler)
}

func unmarshalJsonBytes(content []byte, v interface{}, unmarshaler *Unmarshaler) error {
	var m map[string]interface{}
	if err := jsonz.Unmarshal(content, &m); err != nil {
		return err
	}

	return unmarshaler.Unmarshal(m, v)
}

func unmarshalJsonReader(reader io.Reader, v interface{}, unmarshaler *Unmarshaler) error {
	var m map[string]interface{}
	if err := jsonz.UnmarshalFromReader(reader, &m); err != nil {
		return err
	}

	return unmarshaler.Unmarshal(m, v)
}
