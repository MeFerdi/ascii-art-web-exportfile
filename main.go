package main

import (
	"fmt"
	"net/http"

	"ascii-art/handler"
)

func main() {
	// Serve static files (HTML/CSS) from the "templates" directory
	fs := http.FileServer(http.Dir("templates"))

	http.Handle("/templates/", http.StripPrefix("/templates/", fs))

	// Handle incoming requests to the root path and the /ascii-art path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			handler.HomeHandler(w, r)
		} else if r.URL.Path == "/ascii-art" {
			handler.AsciiArtHandler(w, r)
		} else if r.URL.Path == "/download"{
			handler.ExportAsciiArtHandler(w, r)
		}else {
			http.ServeFile(w, r, "templates/404.html")
		}
	})

	port := ":8080"
	fmt.Printf("Starting Server at port %v\nAt http://localhost:8080", port)

	http.ListenAndServe(port, nil)
}
