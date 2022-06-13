package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/pflag"
)

type LogMessage struct {
	Path    string      `json:"path"`
	Body    string      `json:"body"`
	Headers http.Header `json:"headers"`
}

func main() {
	var showBody, logToStderr bool
	port := 8080
	responseCode := 200
	pflag.BoolVar(&showBody, "show-body", false, "Set if you want the request bodies to be included in log output")
	pflag.BoolVar(&logToStderr, "log-to-stderr", false, "Set if you want logging to go to stderr. It's usually easier to keep it unset if you want to do something like pipe to jq.")
	pflag.IntVar(&port, "port", 8080, "Port to listen on")
	pflag.IntVar(&responseCode, "response-code", 200, "HTTP status code for response, default is 200")
	pflag.Parse()
	if !logToStderr {
		log.SetOutput(os.Stdout)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m := LogMessage{
			Path:    r.URL.Path,
			Headers: r.Header,
		}
		if showBody {
			b, err := io.ReadAll(r.Body)
			if err != nil {
				log.Print("{\"error\":\"Unable to read request body.\"}")
			} else {
				m.Body = string(b)
			}
		}
		err := json.NewEncoder(log.Writer()).Encode(m)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Somehow failed writing logs")
			return
		}
		w.WriteHeader(responseCode)
		fmt.Fprint(w, "{\"result\": \"ok\"")
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
