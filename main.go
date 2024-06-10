package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Default port if PORT environment variable is not set
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pid := fmt.Sprintf("Hello world! Process ID : %d, Port: %s", os.Getpid(), port)
		w.Write([]byte(pid))
	})

	http.ListenAndServe(":"+port, nil)
}
