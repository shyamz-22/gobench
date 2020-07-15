package gobench

import (
	"bytes"
	"encoding/json"
	"sync"
)

type Configuration struct {
	Patterns []string `json:"patterns"`
	Url      string   `json:"url"`
}

var jsonBufferPool = sync.Pool{New: func() interface{} {
	return new(bytes.Buffer)
}}

func JsonMarshal(val Configuration) (string, error) {
	jsonB, err := json.Marshal(&val)

	if err != nil {
		return "", err
	}

	return string(jsonB), nil
}

func JsonEncoder(val Configuration) (string, error) {
	jsonBuffer := jsonBufferPool.Get().(*bytes.Buffer)
	defer func() {
		jsonBuffer.Reset()
		jsonBufferPool.Put(jsonBuffer)
	}()

	encoder := json.NewEncoder(jsonBuffer)
	err := encoder.Encode(val)

	if err != nil {
		return "", err
	}

	return jsonBuffer.String(), nil
}
