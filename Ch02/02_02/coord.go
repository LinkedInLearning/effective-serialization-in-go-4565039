package coord

import (
	"fmt"
	"math"
)

type Coord struct {
	Lat float64
	Lng float64
}

// MarshalText implement encoding.TextMarshaler
func (c Coord) MarshalText() ([]byte, error) {
	lat, lng := math.Abs(c.Lat), math.Abs(c.Lng)
	n := "N"
	if c.Lat < 0 {
		n = "S"
	}

	e := "E"
	if c.Lng < 0 {
		e = "W"
	}

	s := fmt.Sprintf("%.6f%s %.6f%s", lat, n, lng, e)
	return []byte(s), nil
}
