package main

import (
	"database/sql"
	_ "embed"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

type Item struct {
	SKU   string
	Name  string
	Price float64
}

var store = []Item{
	{"m183x", "Magic Wand", 7.0},
	{"m184y", "Invisibility Cape", 13.2},
	{"m185z", "Levitation Spell", 9.3},
}

var (
	//go:embed insert.sql
	insertSQL string

	//go:embed get.sql
	getSQL string
)

func main() {
	db, err := sql.Open("sqlite", "store.db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// Save
	for _, item := range store {
		if _, err := db.Exec(insertSQL, item.SKU, item.Name, item.Price); err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
	}

	// Load
	const sku = "m184y"
	row := db.QueryRow(getSQL, sku)
	if row.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	i := Item{
		SKU: sku,
	}
	if err := row.Scan(&i.Name, &i.Price); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", i)
}
