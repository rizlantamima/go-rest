package main

import (
	"log"
	"net/http"
)

const PORT string = ":8000"

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("hello"))
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting a server...")

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}
