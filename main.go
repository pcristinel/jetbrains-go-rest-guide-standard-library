package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = 8080

type HomeHandler struct {
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", HomeHandler{})

	// Run the server
	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), mux)
	if err != nil {
		log.Panic(err)
	}

	log.Println(fmt.Sprintf("Listening on port %v", PORT))
}

func (h HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("This is my home page"))
	if err != nil {
		log.Println("Error while writing to request: ", err)
		return
	}
}
