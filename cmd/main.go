package main

import (
	"encoding/json"
	"net/http"
)

func main() {

	handler := http.NewServeMux()

	handler.HandleFunc("/", Hello)

	http.ListenAndServe(":8080", handler)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(map[string]string{"data": "Hello World!"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
