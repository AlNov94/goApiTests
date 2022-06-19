package encoders

import "encoding/xml"

type xmlEncoder struct{}

func (xmlEncoder) Encode(v any) ([]byte, error) {
	return xml.Marshal(v)
}

func (xmlEncoder) ToOutput(v any, prefix, indent string) ([]byte, error) {
	return xml.MarshalIndent(v, prefix, indent)
}
