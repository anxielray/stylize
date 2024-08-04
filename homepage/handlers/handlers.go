package handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"artweb/asciiart"
)

type PageData struct {
	Result string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		serveErrorPage(w, r, http.StatusNotFound)
		return
	}

	// Render the form
	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		serveErrorPage(w, r, http.StatusInternalServerError)
		return
	}

	// Create an instance of PageData with the desired result
	data := PageData{
		Result: "", // Replace with your actual result(Ascii-art)
	}

	// Execute the template with the provided data
	if err := tpl.Execute(w, data); err != nil {
		serveErrorPage(w, r, http.StatusInternalServerError)
		return
	}
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		serveErrorPage(w, r, http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/ascii-art" {
		serveErrorPage(w, r, http.StatusNotFound)
	} else {
		// Process the form and generate ASCII art
		input := r.FormValue("message")
		input = strings.ReplaceAll(input, "\r", "\n")
		banner := r.FormValue("banner")
		if input == "" || banner == "" {
			serveErrorPage(w, r, http.StatusBadRequest)
			return
		}

		asciiArt, err := asciiart.Art(input, banner)
		if err != nil {
			if asciiArt == "400" {
				serveErrorPage(w, r, http.StatusBadRequest)
				return
			} else if asciiArt == "404" {
				serveErrorPage(w, r, http.StatusNotFound)
				return
			}
		}

		data := PageData{
			Result: asciiArt,
		}

		w.WriteHeader(http.StatusOK)
		log.Println("200 OK: ASCII art generated successfully")

		_, eer := os.Stat("templates/index.html")
		if eer != nil {
			serveErrorPage(w, r, http.StatusInternalServerError)
			return
		}

		// Render the template with the generated ASCII art
		tpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			serveErrorPage(w, r, http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, data)
	}
}

func serveErrorPage(w http.ResponseWriter, r *http.Request, statusCode int) {
	var filename string

	switch statusCode {
	case http.StatusNotFound:
		filename = "templates/404.html"
	case http.StatusInternalServerError:
		filename = "templates/500.html"
	case http.StatusBadRequest:
		filename = "templates/400.html"
	case http.StatusMethodNotAllowed:
		filename = "templates/405.html" // Assuming you have a 405.html template
	default:
		filename = "templates/500.html" // Fallback to 500
	}

	// Load the HTML template
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the status code and execute the template
	w.WriteHeader(statusCode)
	tmpl.Execute(w, nil)
}
