package decoder

import (
	"bytes"
	"encoding/json"
)

type JsonNumberDecoder struct{}

func (cd *JsonNumberDecoder) Decode(data []byte, v interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	return decoder.Decode(v)
}

func (cd *JsonNumberDecoder) Encode(v interface{}) ([]byte, error) {
	b := new(bytes.Buffer)
	encoder := json.NewEncoder(b)
	err := encoder.Encode(v)
	return b.Bytes(), err
}

func JosonNumberEncode(v interface{}) ([]byte, error) {
	var cd JsonNumberDecoder

	return cd.Encode(v)
}
func JsonNumberDecode(data []byte, v interface{}) error {
	var cd JsonNumberDecoder
	err := cd.Decode(data, v)
	if err != nil {
		return err
	}

	return nil
}
