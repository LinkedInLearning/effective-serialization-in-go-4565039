package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("store.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	r := csv.NewReader(file)
	for {
		fields, err := r.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}

		fmt.Println(fields)
	}
}
