package main

import (
	"github.com/TerryHowe/goat/v1"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", VersionGet).Methods("GET")
	r.HandleFunc("/v1", VersionGet).Methods("GET")
	r.HandleFunc("/v1/clusters", v1.ClusterGet).Methods("GET")
	http.Handle("/", r)

	// wait for clients
	http.ListenAndServe(":8083", nil)
}
