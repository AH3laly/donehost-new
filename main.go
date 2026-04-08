package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Donehost server!")
}

func main() {
	http.HandleFunc("/", handler)

	port := "8601"
	fmt.Println("Server running on port", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
