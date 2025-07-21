package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(
		w,
		`{
			"message": "Hello, World!"
		}`,
	)
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(
		w,
		`{
			"message": "Welcome to our API!"
		}`,
	)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/welcome", welcomeHandler)

	port := ":80"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
