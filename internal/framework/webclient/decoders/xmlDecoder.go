package decoders

import "encoding/xml"

type XmlDecoder struct{}

func (XmlDecoder) Decode(data []byte, v any) error {
	return xml.Unmarshal(data, v)
}
