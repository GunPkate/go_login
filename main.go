package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func JSON(w http.ResponseWriter, statusCode int) func(v any) {

	w.Header().Set("Content-Type", "application/jsosn: character=UTF8")
	w.WriteHeader(statusCode)

	return func(v any) {
		json.NewEncoder(w).Encode(v)
	}

}

func main() {
	r := mux.NewRouter()
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hello, World")
	// })
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]any{
			"message": "Hello world",
		}
		JSON(w, http.StatusOK)(resp)

		// json.NewEncoder(w).Encode(resp)
		// fmt.Fprint(w, "Hello, World")
	}).Methods(http.MethodGet)
	// }).Methods("GET")
	http.ListenAndServe(":8000", r)
}
