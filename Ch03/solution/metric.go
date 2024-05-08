package metrics

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	var buf bytes.Buffer
	fmt.Fprintf(&buf, m.Name)
	if len(m.Labels) > 0 {
		fmt.Fprintf(&buf, "{")
		i := 0
		for k, v := range m.Labels {
			fmt.Fprintf(&buf, "%s=%q", k, v)
			if i < len(m.Labels)-1 {
				fmt.Fprint(&buf, ",")
			}
			i++
		}
		fmt.Fprintf(&buf, "}")
	}
	fmt.Fprintf(&buf, " %f", m.Value)
	if !m.Time.IsZero() {
		fmt.Fprintf(&buf, " %d", m.Time.UnixMicro())
	}

	return buf.Bytes(), nil
}

type Metrics []Metric

func (ms Metrics) MarshalJSON() ([]byte, error) {
	m := map[string]any{
		"count":   len(ms),
		"metrics": []Metric(ms),
	}

	return json.Marshal(m)
}
