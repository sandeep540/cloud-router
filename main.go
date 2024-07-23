package main

import (
	"cloud-router/bootstrap"
	"log/slog"
	"net/http"
	"os"
)

func main() {

	app := bootstrap.App()
	env := app.Env

	port := env.ServerAddress

	v1 := http.NewServeMux()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	v1.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	v1.HandleFunc("/", handler)

	v1.HandleFunc("/v1/yahoo", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://yahoo.com", http.StatusSeeOther)
	})

	logger.Info("Spinning Server on Port " + port)
	http.ListenAndServe(port, v1)

}

func handler(w http.ResponseWriter, r *http.Request) {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	logger.Info("Method: ", r.Method, " URL: ", r.RequestURI, " Proto: ", r.Proto)

	//Iterate over all header fields
	/* for k, v := range r.Header {
		logger.Info("Header field: ", k, " Value: ", v)
	} */

	// Read all the headers
	for name, headers := range r.Header {
		// Iterate all headers with one name (e.g. Content-Type)
		for _, hdr := range headers {
			//println(name + ": " + hdr)
			logger.Info("Header name: ", name, " Value: ", hdr)
		}
	}

	logger.Info("Host = ", r.Host)
	logger.Info("RemoteAddr= ", r.RemoteAddr)
	//Get value for a specified token
	logger.Info("Finding value of Accept", r.Header["Accept"])

	w.Write([]byte("Test Success"))
}
