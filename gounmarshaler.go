package json

import "github.com/go-summer-dev/marshaling"

type gounmarshaler struct {
	marshaling.Unmarshaler
}

func (g *gounmarshaler) UnmarshalJSON(b []byte) error {
	return g.Unmarshaler.GoUnmarshal(b)
}
