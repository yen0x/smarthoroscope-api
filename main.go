package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"./horoscope"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/stats", horoscope.GetStats)
	fmt.Println("Started...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "SmartHoroscope API")
}
