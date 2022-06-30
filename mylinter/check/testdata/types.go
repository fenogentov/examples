package testdata

type Document struct {
	m    map[string]any
	keys []string
}

type Array struct {
	s []any
}

type Binary struct {
	Subtype string
	B       []byte
}

type ObjectID map[string]byte

type NullType struct{}

type Regex struct {
	Pattern string
	Options string
}

type Timestamp int64
