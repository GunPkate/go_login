package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hello, World")
	// })
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/jsosn: character=UTF8")
		w.WriteHeader(200)

		resp := map[string]any{
			"message": "Hello world",
		}

		json.NewEncoder(w).Encode(resp)
		// fmt.Fprint(w, "Hello, World")
	}).Methods(http.MethodGet)
	// }).Methods("GET")
	http.ListenAndServe(":8000", r)
}
