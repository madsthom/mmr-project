package main

import (
	"fmt"
	"net/http"

	database "example.com/m/v2/db"
)

func main() {
	// Initialize Database
	database.InitDatabase()
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
