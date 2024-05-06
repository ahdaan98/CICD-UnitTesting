package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/enescakir/emoji"
)

const PORT = ":3000"

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	message := fmt.Sprintf("Hello, world! %v", emoji.Popcorn)
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", handler)

	log.Println("Server Running on PORT:",PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}