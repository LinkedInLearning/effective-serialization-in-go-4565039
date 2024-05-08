package date

import (
	"encoding/json"
	"fmt"
	"time"
)

func ExampleDate() {
	d := Date{2024, time.April, 2}
	data, err := json.Marshal(d)
	if err != nil {
		fmt.Println("error (marshal):", err)
		return
	}

	fmt.Println(string(data))

	var d2 Date
	if err := json.Unmarshal(data, &d2); err != nil {
		fmt.Println("error (unmarshal):", err)
		return
	}
	fmt.Println(d2)

	// Output:
	// "2024-04-02"
	// {2024 April 2}
}
