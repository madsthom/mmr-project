package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/greeting", greetingHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing name parameter", http.StatusBadRequest)
		return
	}

	message := fmt.Sprintf("Hello, %s!", name)
	fmt.Fprintln(w, message)
}
