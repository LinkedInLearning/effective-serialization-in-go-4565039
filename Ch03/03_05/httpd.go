package main

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

func logsHandler(w http.ResponseWriter, r *http.Request) {
	ctrl := http.NewResponseController(w)
	rdr, wtr := io.Pipe()
	w.Header().Set("Content-Type", "application/jsonl")
	go Producer(wtr, 10)

	s := bufio.NewScanner(rdr)
	for s.Scan() {
		w.Write(s.Bytes())
		w.Write([]byte{'\n'})
		ctrl.Flush()
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /logs", logsHandler)

	addr := ":8080"
	srv := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	slog.Info("server starting", "address", addr)
	if err := srv.ListenAndServe(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
