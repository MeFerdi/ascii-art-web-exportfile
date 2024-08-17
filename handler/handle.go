package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
	"time"
	"ascii-art/ascii"
)

var templates *template.Template

// init initializes the template by parsing the index.html file.
func init() {
	var err error
	templates, err = template.ParseFiles(filepath.Join("templates", "index.html"))
	if err != nil {
		fmt.Println("Unable to parseFile: index.html missing")
		os.Exit(0)
	}
}

// HomeHandler handles requests to the home page.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.ServeFile(w, r, "templates/405.html")
		return
	}

	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template index.html: %v", err)
	}
}

// AsciiArtHandler handles requests for generating ASCII art.
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.ServeFile(w, r, "templates/405.html")
		return
	}

	str := r.FormValue("textData")
	bannerStyle := r.FormValue("banner")

	if len(str) == 0 {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	art, err := ascii.PrintAscii(str, bannerStyle)
	if err != nil {
		http.ServeFile(w, r, "templates/500.html")
		return
	}

	// Prepare the data to pass to the template
	data := struct {
		Art string
	}{
		Art: art,
	}

	renderTemplate(w, "index", data)
}

// renderTemplate renders the specified template with the provided data.
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	if templates == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template %s: %v", tmpl, err)
	}
}

// DownloadAsciiArtHandler handles the downloading of the ASCII art.
func ExportAsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Retrieve the ASCII art from the form
		asciiArt := r.FormValue("asciiArt")

		if len(asciiArt) == 0 {
			http.Error(w, "400 Bad Request: No ASCII art provided", http.StatusBadRequest)
			return
		}

		// Create a temporary file to store the ASCII art
		fileName := fmt.Sprintf("ascii-art-%d.txt", time.Now().UnixNano())
		filePath := filepath.Join(os.TempDir(), fileName)

		err := os.WriteFile(filePath, []byte(asciiArt), 0644)
		if err != nil {
			http.Error(w, "Could not generate file", http.StatusInternalServerError)
			log.Printf("Error writing file: %v", err)
			return
		}
		// Calculate the content length based on the ASCII art string length
		contentLength := len(asciiArt)

		// Serve the file as a download
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", contentLength))
		http.ServeFile(w, r, filePath)

		// Clean up the temporary file after serving
		defer os.Remove(filePath)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
