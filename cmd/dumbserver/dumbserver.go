package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/pflag"
)

type LogMessage struct {
	Path    string      `json:"path"`
	Body    string      `json:"body"`
	Headers http.Header `json:"headers""`
}

func main() {
	var showBody bool
	pflag.BoolVar(&showBody, "show-body", false, "Set to true if you want the request bodies to be included in log output")
	pflag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m := LogMessage{
			Path:    r.URL.Path,
			Headers: r.Header,
		}
		if showBody {
			b, err := io.ReadAll(r.Body)
			if err != nil {
				log.Print("{\"error\":\"Unable to read request body.\"}")
			}
			m.Body = string(b)
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
