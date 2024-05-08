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

type Metrics []Metric
