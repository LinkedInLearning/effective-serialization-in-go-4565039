package main

import (
	"bytes"
	"fmt"
	"os"
)

func ExampleProducer() {
	err := Producer(os.Stdout, 3)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// Output:
	// {"time":"2024-04-03T11:47:35Z","level":"info","message":"log #1"}
	// {"time":"2024-04-03T11:47:35.523Z","level":"info","message":"log #2"}
	// {"time":"2024-04-03T11:47:36.046Z","level":"info","message":"log #3"}
}

func ExampleConsumer() {
	data := []byte(`
		{"time":"2024-04-03T11:47:35Z","level":"info","message":"log #1"}
		{"time":"2024-04-03T11:47:35.523Z","level":"info","message":"log #2"}
		{"time":"2024-04-03T11:47:36.046Z","level":"info","message":"log #3"}
	`)

	r := bytes.NewReader(data)
	err := Consumer(r)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// Output:
	// log: {2024-04-03 11:47:35 +0000 UTC info log #1}
	// log: {2024-04-03 11:47:35.523 +0000 UTC info log #2}
	// log: {2024-04-03 11:47:36.046 +0000 UTC info log #3}
}
