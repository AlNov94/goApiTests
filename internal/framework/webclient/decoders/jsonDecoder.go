package decoders

import "encoding/json"

type JsonDecoder struct{}

func (JsonDecoder) Decode(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
