package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handlePackageHook(w http.ResponseWriter, r *http.Request) {
	// Handle the package hook here
	fmt.Println(r.Method)
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	fmt.Println(data)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Package hook received"))
}

func main() {
	http.HandleFunc("/api/packages", handlePackageHook)
	log.Fatal(http.ListenAndServe(":8001", nil))
}
