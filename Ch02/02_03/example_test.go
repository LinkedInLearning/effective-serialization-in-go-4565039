package store

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"time"
)

func ExampleItem() {
	i := Item{
		SKU:   "BTN-FF0000",
		Name:  "Red Button",
		Price: 75,
		Added: time.Date(2024, time.April, 1, 13, 17, 43, 50, time.UTC),
		Image: nil, // TODO Add image
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(i); err != nil {
		fmt.Println("error (encode):", err)
		return
	}

	dec := gob.NewDecoder(&buf)
	var i2 Item
	if err := dec.Decode(&i2); err != nil {
		fmt.Println("error (decode):", err)
		return
	}

	fmt.Printf("%s: %s\n", i2.SKU, i2.Name)

	// Output:
	// BTN-FF0000: Red Button
}

func ExampleStream() {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	for i := range 3 {
		i := Item{
			SKU:  fmt.Sprintf("SKU-%d", i),
			Name: fmt.Sprintf("Item #%d", i),
		}
		if err := enc.Encode(i); err != nil {
			fmt.Println("error (encode):", err)
			return
		}
	}

	dec := gob.NewDecoder(&buf)
	for {
		var i Item
		err := dec.Decode(&i)

		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Println("error (decode):", err)
			return
		}

		fmt.Printf("%s: %s\n", i.SKU, i.Name)
	}

	// Output:
	// SKU-0: Item #0
	// SKU-1: Item #1
	// SKU-2: Item #2
}
