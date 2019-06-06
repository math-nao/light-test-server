package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", LandingPageAction)

	listener := ":8080"
	log.Println("Listening on " + listener + "...")
	http.ListenAndServe(listener, nil)
}

// LandingPageAction handles the landing page request
func LandingPageAction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome. Date: "+time.Now().Format("2006.01.02 15:04:05 MST"))
}
