package encoders

import "encoding/xml"

type XmlEncoder struct{}

func (XmlEncoder) Encode(v any) ([]byte, error) {
	return xml.Marshal(v)
}

func (XmlEncoder) ToOutput(v any, prefix, indent string) ([]byte, error) {
	return xml.MarshalIndent(v, prefix, indent)
}
