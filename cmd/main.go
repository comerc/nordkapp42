package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	handler := http.NewServeMux()

	handler.HandleFunc("/", Hello)

	// cert, err := tls.LoadX509KeyPair("localhost.pem", "localhost-key.pem")
	// if err != nil {
	// 	log.Fatalf("server: loadkeys: %s", err)
	// }
	// config := tls.Config{Certificates: []tls.Certificate{cert}}

	server := &http.Server{
		Addr:    ":8888",
		Handler: handler,
	}

	log.Printf("Listening to port 8888")
	log.Fatal(server.ListenAndServeTLS("localhost.pem", "localhost-key.pem"))
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
