package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for _, e := range os.Environ() {
			fmt.Fprintf(os.Stdout, "%s\n", e)
			fmt.Fprintf(w, "%s\n", e)
		}
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
