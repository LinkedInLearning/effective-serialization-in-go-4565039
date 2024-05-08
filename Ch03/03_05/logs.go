package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"
)

// Log is a log in the database.
type Log struct {
	Time    time.Time `json:"time"`
	Level   string    `json:"level"`
	Message string    `json:"message"`
}

// Producer agent tailing logs.
func Producer(w io.WriteCloser, n int) error {
	defer w.Close()

	enc := json.NewEncoder(w)
	start := time.Date(2024, time.April, 3, 11, 47, 35, 0, time.UTC)
	delta := 523 * time.Millisecond

	for i := range n {
		log := Log{
			Time:    start.Add(time.Duration(i) * delta),
			Level:   "info",
			Message: fmt.Sprintf("log #%d", i+1),
		}
		if err := enc.Encode(log); err != nil {
			return err
		}
		time.Sleep(100 * time.Millisecond) // Simulate network latency
	}

	return nil
}

// Consumer reads logs from r.
func Consumer(r io.Reader) error {
	dec := json.NewDecoder(r)

	for {
		var log Log
		err := dec.Decode(&log)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return err
		}

		fmt.Println("log:", log) // TODO: Handle logs
	}

	return nil
}
