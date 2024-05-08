package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Version int
	Server  struct {
		Port int
	}
	Logging struct {
		Level string
	}
}

func main() {
	file, err := os.Open("config.yml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var cfg Config
	dec := yaml.NewDecoder(file)
	if err := dec.Decode(&cfg); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", cfg)
}
