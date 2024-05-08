package metrics

import (
	"fmt"
	"time"
)

func ExampleMetric_MarshalText() {
	// Full data
	m := Metric{
		Time:  time.Date(2024, time.February, 7, 4, 32, 56, 392, time.UTC),
		Name:  "http_errors",
		Value: 3,
		Labels: map[string]string{
			"host": "srv1",
			"app":  "reports",
		},
	}

	data, err := m.MarshalText()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(data))

	// No Labels
	m = Metric{
		Time:  time.Date(2024, time.February, 7, 4, 32, 56, 392, time.UTC),
		Name:  "http_errors",
		Value: 3,
	}

	data, err = m.MarshalText()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(data))

	// No time
	m = Metric{
		Name:  "http_errors",
		Value: 3,
		Labels: map[string]string{
			"host": "srv1",
			"app":  "reports",
		},
	}

	data, err = m.MarshalText()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(data))

	// No time & labels
	m = Metric{
		Name:  "http_errors",
		Value: 3,
	}

	data, err = m.MarshalText()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(data))

	// Output:
	// http_errors{host="srv1",app="reports"} 3.000000 1707280376000000
	// http_errors 3.000000 1707280376000000
	// http_errors{host="srv1",app="reports"} 3.000000
	// http_errors 3.000000
}
