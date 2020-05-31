package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/victoralagwu/gophercises/urlshort/handler"
	"github.com/victoralagwu/gophercises/urlshort/mysqlservice/models"
)

// Server : Implement
type Server struct {
	DB *gorm.DB
	Router *mux.Router

}
var pathsToUrls = map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
}


func main() {
	// Use flags to indicate either to use YAML or JSON
	// selectedType := flag.String("type", "yaml", "Please select the right file to use, yaml/json")
	// flag.Parse()

	mux := defaultMux()
	
	// Build the MapHandler using the mux as the fallback

	fallback := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// // fallback

	var yamlhandler = createYAMLHandler(fallback)
	var jsonHandler = createJSONHandler(yamlhandler)
	fmt.Println("Starting the server on :8085")
	http.ListenAndServe(":8085", jsonHandler)
}

//Initialize :
func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort,DbHost, DbName string) {
	var err error
	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		server.DB, err = gorm.Open(Dbdriver, DBURL)

		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("Database Connection error: ", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	server.DB.Debug().AutoMigrate(&models.Link{})
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



func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}