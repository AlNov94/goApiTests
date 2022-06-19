package encoders

type Encoder interface {
	Encode(v any) ([]byte, error)
	ToOutput(v any, prefix, indent string) ([]byte, error)
}
