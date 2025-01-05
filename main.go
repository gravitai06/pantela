package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var task string

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/task", TasksHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", task)
}

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Task string `json:"task"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	task = request.Task
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Task updated")
}
