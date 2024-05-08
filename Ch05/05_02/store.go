package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Store struct {
	Items []struct {
		SKU   string  `xml:"sku,attr"`
		Name  string  `xml:"name"`
		Price float64 `xml:"price"`
	} `xml:"item"`
}

func main() {
	file, err := os.Open("store.xml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var s Store
	dec := xml.NewDecoder(file)
	if err := dec.Decode(&s); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", s)
}
