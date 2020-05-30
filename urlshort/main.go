package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"flag"

	"github.com/victoralagwu/gophercises/urlshort/handler"
)
func main() {
	// Use flags to indicate either to use YAML or JSON
	selectedType := flag.String("type", "yaml", "Please select the right file to use, yaml/json")
	flag.Parse()

	mux := defaultMux()
	
	// Build the MapHandler using the mux as the fallback

	fallback := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	var handler http.HandlerFunc
	if *selectedType == "json" {
		fmt.Printf("Default Selected Type: %s\n", *selectedType)
		handler = createJSONHandler(fallback)
	} else {
		fmt.Printf("Default Selected Type: %s\n", *selectedType)
		handler = createYAMLHandler(fallback)
	}

	fmt.Println("Starting the server on :8085")
	http.ListenAndServe(":8085", handler)
}
func createYAMLHandler(fallback http.HandlerFunc) http.HandlerFunc {	
	yamlFile, err := ioutil.ReadFile("urls.yml")
	if err != nil {
		panic(err)
	}
	yamlHandler, err := urlshort.YAMLHandler([]byte(yamlFile), fallback)
	if err != nil {
		panic(err)
	}
	return yamlHandler
}

func createJSONHandler(fallback http.HandlerFunc) http.HandlerFunc {
	//Read Json File
	jsonFile, err := ioutil.ReadFile("urls.json")
	if err != nil {
		panic(err)
	}
	jsonHandler, err := urlshort.JSONHandler([]byte(jsonFile), fallback)
	if err != nil {
		panic(err)
	}
	return jsonHandler
}

var pathsToUrls = map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}