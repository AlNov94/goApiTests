package decoders

import "encoding/xml"

type xmlDecoder struct{}

func (xmlDecoder) Decode(data []byte, v any) error {
	return xml.Unmarshal(data, v)
}
