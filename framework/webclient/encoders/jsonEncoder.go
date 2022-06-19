package encoders

import "encoding/json"

type JsonEncoder struct{}

func (JsonEncoder) Encode(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (JsonEncoder) ToOutput(v any, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}
