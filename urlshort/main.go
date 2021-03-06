package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"flag"
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql datatype driver
	"github.com/victoralagwu/gophercises/urlshort/handler"
	"github.com/victoralagwu/gophercises/urlshort/mysqlservice/models"
	"github.com/victoralagwu/gophercises/urlshort/mysqlservice/seed"
)

// Server : Implement
type Server struct {
	DB *gorm.DB
}
var pathsToUrls = map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
}

var server = Server{};

const (
	seederFlagValue = false
	seederUsage     = "Usage: -seeder=true or -seeder=false"
	port = ":8085"
)

func main() {
	var err error
	// Use flags to indicate either to use YAML or JSON
	// selectedType := flag.String("type", "yaml", "Please select the right file to use, yaml/json")
	useSeeder := flag.Bool("seeder", seederFlagValue, seederUsage)
	
	flag.Parse()

	mux := defaultMux()
	
	// Build the MapHandler using the mux as the fallback

	fallback := urlshort.MapHandler(pathsToUrls, mux)

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not coming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}


	server.Initialize(os.Getenv("DB_DRIVER"),
							os.Getenv("DB_USER"),
							os.Getenv("DB_PASSWORD"),
							os.Getenv("DB_PORT"),
							os.Getenv("DB_HOST"),
							os.Getenv("DB_NAME"))

	//See data
	if *useSeeder == true {
		fmt.Print("\nSeeder loaded")
		seed.Load(server.DB)
	} else {
		fmt.Print("\nSeeder not loaded")
	}

		// Build the YAMLHandler using the mapHandler as the
	// // fallback

	var yamlhandler = createYAMLHandler(fallback)
	var jsonHandler = createJSONHandler(yamlhandler)
	var mysqlHandler = createMYSQLHandler(jsonHandler)

	fmt.Printf("Starting the server on %s\n", port)
	http.ListenAndServe(port, mysqlHandler)
}

//Initialize :
func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort,DbHost, DbName string) {
	var err error

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	server.DB, err = gorm.Open(Dbdriver, DBURL)

	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", Dbdriver)
		log.Fatal("Database Connection error: ", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}
		// defer server.DB.Close()
	

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

func createMYSQLHandler(fallback http.HandlerFunc) http.HandlerFunc {
	//Get data from database
	link := models.Link{}
	links, err := link.FindAll(server.DB)
	if err != nil {
		panic(err)
	}
	j, err := json.Marshal(links)
	if err != nil {
		panic(err)
	}
	mySQLHandler, err := urlshort.SQLHandler([]byte(j), fallback)
	if err != nil {
		panic(err)
	}
	return mySQLHandler
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func loadHello(w http.ResponseWriter, r *http.Request, data *http.HandlerFunc) {
	fmt.Fprintln(w, data)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "URL Shortner")
}