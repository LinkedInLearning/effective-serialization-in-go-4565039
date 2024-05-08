package date

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

type Date struct {
	Year  int
	Month time.Month
	Day   int
}

func (d Date) MarshalJSON() ([]byte, error) {
	// Step 1: Convert to type json.Marshal can handle
	s := fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)

	// Step 2: Use json.Marshal
	return json.Marshal(s)

	// Step 3: There is no step 3 ☺
}

func (d *Date) UnmarshalJSON(data []byte) error {
	var year, month, day int

	r := bytes.NewReader(data)
	if _, err := fmt.Fscanf(r, `"%04d-%02d-%02d"`, &year, &month, &day); err != nil {
		return err
	}

	d.Year, d.Month, d.Day = year, time.Month(month), day
	return nil
}
