package decoders

import "encoding/json"

type jsonDecoder struct{}

func (jsonDecoder) Decode(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
