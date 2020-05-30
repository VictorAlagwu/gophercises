package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/victoralagwu/gophercises/urlshort"
)
func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	// yamlFile, err := ioutil.ReadFile("urls.yml")
	// if err != nil {
	// 	panic(err)
	// }
	// yamlHandler, err := urlshort.YAMLHandler([]byte(yamlFile), mapHandler)
	// if err != nil {
	// 	panic(err)
	// }

	//Read Json File
	jsonFile, err := ioutil.ReadFile("urls.json")
	if err != nil {
		panic(err)
	}
	jsonHandler, err := urlshort.JSONHandler([]byte(jsonFile), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8085")
	http.ListenAndServe(":8085", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}