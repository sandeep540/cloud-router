package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {

	v1 := http.NewServeMux()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	v1.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	v1.HandleFunc("/v1/yahoo", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://yahoo.com", http.StatusSeeOther)
	})

	logger.Info("Spinning Server on Port 8000")
	http.ListenAndServe(":8000", v1)

}
