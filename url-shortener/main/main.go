package main

import (
	"fmt"
	"gobasicprojects/urlshortener"
	"net/http"
	"os"
)

func main() {
	// Initialize the default multiplexer (mux) with a basic hello handler
	mux := defaultMux()

	// Define path mappings with URLs
	pathsToUrls := map[string]string{
		"/doc": "https://go.dev/doc",
		"/pkg": "https://pkg.go.dev",
	}

	// Create the MapHandler that checks pathsToUrls first, falling back to mux
	mapHandler := urlshortener.MapHandler(pathsToUrls, mux)

	// YAML configuration for additional URL mappings
	yaml := `
- path: /learn
  url: https://go.dev/learn/
`

	// Parse the YAML and create a YAMLHandler with mapHandler as fallback
	yamlHandler, err := urlshortener.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		exit(fmt.Sprintf("Failed to parse yaml: %s\n", err))
	}

	// Start the server with error handling
	port := ":8080"
	fmt.Printf("Starting the server on %s\n", port)
	if err := http.ListenAndServe(port, yamlHandler); err != nil {
		exit(fmt.Sprintf("Server failed to start: %s\n", err))
	}
}

// defaultMux initializes the default multiplexer with a root handler
func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello) // Handle the root path with hello handler
	return mux
}

// hello is a default handler for the root path
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

// exit prints an error message and exits the program
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
