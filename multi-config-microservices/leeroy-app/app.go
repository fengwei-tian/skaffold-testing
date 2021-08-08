package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "leeroooooy app - let's go modus!!\n")
}

func main() {
	log.Print("leeroy app server ready")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":50051", nil)
}
