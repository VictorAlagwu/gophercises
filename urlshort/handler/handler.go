package urlshort

import (
	"encoding/json"
	"fmt"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)
type pathURL struct {
	Path string `yaml:"path" json:"path"`
	URL string	`yaml:"url" json:"url"`
}
// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	//Map through pathsToUrls
	return func(w http.ResponseWriter, r *http.Request) {
		currentPath := r.URL.Path
		if i, status := pathsToUrls[currentPath]; status {
			fmt.Printf("Req: %s \n", i) 
			http.Redirect(w, r, i , http.StatusFound)
			return
		} 
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	pathUrls, err := parseYaml(yml)
	if err != nil {
		return nil, err
	}
	//Put parse yml in map of string [{key: value}]
	pathMap := buildMap(pathUrls)
	return MapHandler(pathMap, fallback), nil
}

func parseYaml(data []byte) ([]pathURL, error){
	var pathUrls []pathURL
	//Parse YAML file
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

func parseJSON(data []byte) ([]pathURL, error){
	var pathUrls []pathURL
	//Parse YAML file
	err := json.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}
// JSONHandler : Returns eqivalent of MapHandler from Json file
func JSONHandler(j []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseJSON(j)
	
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(pathUrls)
	return MapHandler(pathMap, fallback), nil
}

//SQLHandler : Implement
func SQLHandler(data []byte, fallback http.Handler) (http.HandlerFunc, error) {
	//Connect to database
	//Get and Parse data from database
	// Build Map from parse data
	return MapHandler([]byte, fallback), nil
}
func buildMap(pathUrls []pathURL) map[string]string {
	pathsToUrls := make(map[string]string) 
	for _, p := range pathUrls {
		pathsToUrls[p.Path] = p.URL
	}
	return pathsToUrls
}