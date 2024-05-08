package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Task struct {
	User     string `json:"user"`
	Priority int    `json:"priority"`
	Payload  []byte `json:"payload"`
}

func main() {
	t := Task{
		User:     "garfield",
		Priority: 1,
		Payload:  []byte("nap"),
	}

	// []byte API
	data, err := json.Marshal(t)
	if err != nil {
		fmt.Println("error (marshal):", err)
		return
	}
	fmt.Println("data:", string(data))

	var t2 Task
	if err := json.Unmarshal(data, &t2); err != nil {
		fmt.Println("error (unmarshal):", err)
		return
	}

	// io.Writer/io.Reader API
	var buf bytes.Buffer

	enc := json.NewEncoder(&buf)
	if err := enc.Encode(t); err != nil {
		fmt.Println("error (encode):", err)
		return
	}
	fmt.Println("data:", buf.String())

	dec := json.NewDecoder(&buf)
	if err := dec.Decode(&t2); err != nil {
		fmt.Println("error (decode):", err)
		return
	}
}
