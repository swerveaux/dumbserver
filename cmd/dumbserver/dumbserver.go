package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type LogMessage struct {
	Path    string      `json:"path"`
	Headers http.Header `json:"headers""`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m := LogMessage{
			Path:    r.URL.Path,
			Headers: r.Header,
		}
		err := json.NewEncoder(log.Writer()).Encode(m)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Somehow failed writing logs")
			return
		}
		fmt.Fprint(w, "OK")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
