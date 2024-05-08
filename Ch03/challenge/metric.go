package metrics

import (
	"time"
)

type Metric struct {
	Time   time.Time
	Name   string
	Value  float64
	Labels map[string]string
}

// MarshalText implements encoding.TextMarshaler
func (m Metric) MarshalText() ([]byte, error) {
	// TODO
	return nil, nil
}
