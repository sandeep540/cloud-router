package main

import (
	"cloud-router/bootstrap"
	"log/slog"
	"net/http"
	"os"
	"regexp"
)

/**
* Various Regex -  https://regex-generator.olafneumann.org/
 */
var rNum = regexp.MustCompile(`\d`)            // Has digit(s)
var rAbc = regexp.MustCompile(`abc`)           // Contains "abc"
var rSem = regexp.MustCompile(`(?i)1\\.0\\.x`) //semVar 1.0.x
var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func main() {

	app := bootstrap.App()
	env := app.Env

	v1 := http.NewServeMux()

	v1.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	v1.HandleFunc("/v1/yahoo", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://yahoo.com", http.StatusSeeOther)
	})

	v1.HandleFunc("/", handler)

	v1.HandleFunc("/explorer/", route)

	port := env.ServerAddress
	logger.Info("Spinning Server on Port " + port)
	http.ListenAndServe(port, v1)

}

func route(w http.ResponseWriter, r *http.Request) {
	switch {
	case rNum.MatchString(r.URL.Path):
		digits(w, r)
	case rAbc.MatchString(r.URL.Path):
		abc(w, r)
	case rSem.MatchString(r.URL.Path):
		sem(w, r)
	default:
		w.Write([]byte("Unknown Pattern"))
	}
}

func digits(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Has digits"))
}

func abc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Has abc"))
}

func sem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Has semVar"))
}

func handler(w http.ResponseWriter, r *http.Request) {

	logger.Info("Printing", "Method: ", r.Method, " URL: ", r.RequestURI, " Proto: ", r.Proto)

	// Read all the headers
	for name, headers := range r.Header {
		for _, hdr := range headers {
			logger.Info("Header name: ", name, " Value: ", hdr, "")
		}
	}

	logger.Info("Host = ", r.Host, "")
	logger.Info("RemoteAddr= ", r.RemoteAddr, "")

	//Get value for a specified token
	logger.Info("Finding value of Accept", r.Header["Accept"])

	w.Write([]byte("Test Success"))
}
