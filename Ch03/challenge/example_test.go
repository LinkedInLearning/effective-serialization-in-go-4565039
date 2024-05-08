package metrics

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func ExampleMetrics_MarshalJSON() {
	// Full data
	m1 := Metric{
		Time:  time.Date(2024, time.February, 7, 4, 32, 56, 392, time.UTC),
		Name:  "http_errors",
		Value: 3,
		Labels: map[string]string{
			"host": "srv1",
			"app":  "reports",
		},
	}

	// No Labels
	m2 := Metric{
		Time:  time.Date(2024, time.February, 7, 4, 32, 56, 392, time.UTC),
		Name:  "http_errors",
		Value: 3,
	}

	// No time
	m3 := Metric{
		Name:  "http_errors",
		Value: 3,
		Labels: map[string]string{
			"host": "srv1",
			"app":  "reports",
		},
	}

	// No time & labels
	m4 := Metric{
		Name:  "http_errors",
		Value: 3,
	}

	ms := Metrics([]Metric{m1, m2, m3, m4})
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("  ", "  ")
	if err := enc.Encode(ms); err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	// Output:
	// {
	//     "count": 4,
	//     "metrics": [
	//       "http_errors{host=\"srv1\",app=\"reports\"} 3.000000 1707280376000000",
	//       "http_errors 3.000000 1707280376000000",
	//       "http_errors{host=\"srv1\",app=\"reports\"} 3.000000",
	//       "http_errors 3.000000"
	//     ]
	//   }
}
