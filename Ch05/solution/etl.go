/*
Insert all the log entries from the XML files in the "logs" directory into logs.db

Create the database:

	sqlite3 logs.db < schema.sql

Run your code.

There should be a total of 991 logs

	echo 'SELECT COUNT(time) FROM logs' | sqlite3 logs.db
*/
package main

import (
	"database/sql"
	_ "embed"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	_ "modernc.org/sqlite"
)

//go:embed insert.sql
var insertSQL string

func insert(db *sql.DB, r io.Reader) (int, error) {
	var doc struct {
		Logs []struct {
			Time    time.Time `xml:"time"`
			Level   string    `xml:"level"`
			Message string    `xml:"message"`
		} `xml:"log"`
	}

	dec := xml.NewDecoder(r)
	if err := dec.Decode(&doc); err != nil {
		return 0, err
	}

	for _, log := range doc.Logs {
		if _, err := db.Exec(insertSQL, log.Time, log.Level, log.Message); err != nil {
			return 0, err
		}
	}

	return len(doc.Logs), nil
}

func main() {
	matches, err := filepath.Glob("logs/*.xml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	db, err := sql.Open("sqlite", "logs.db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	defer db.Close()

	for _, fileName := range matches {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
		defer file.Close()

		n, err := insert(db, file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s: %d records\n", fileName, n)
	}
}
