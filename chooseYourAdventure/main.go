package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"

	// "io/ioutil"
	"log"
	"net/http"
	"os"
)
//Story : 
type Story map[string]Chapter 

//Chapter :
type Chapter struct{
	Title   string   `json:"title"`
	Paragraphs   []string `json:"story"`
	Options []Option `json:"options"`
}

//Option :
type Option struct {
	Text string `json:"text"`
	Chapter  string `json:"arc"`
}

func startApp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Choose Your Adventure: " + r.URL.Path[1:])
}

func main() {
	//
	jsonFile, err := createJSONHandler()
	if err != nil {
		fmt.Println(err)
	}

	//Start HTTP Server
	fmt.Println("Server Started on Port 8082")
	http.HandleFunc("/view/", viewHandler)
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		jData, err := json.Marshal(jsonFile)
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("Server started")
		w.Write(jData)
	})
	log.Fatal(http.ListenAndServe(":8082", nil))
}



func loadPage(title string) (Story, error) {
    // filename := title + ".txt"
	// body, err := ioutil.ReadFile(filename)
    // if err != nil {
    //     return nil, err
	// }
	jsonFile, err := createJSONHandler()
	if err != nil {
		fmt.Println(err)
	}
    return jsonFile, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	s, err := loadPage("ds")
	if err != nil {
		fmt.Println(err)
	}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, s)
}

func createJSONHandler() (Story, error) {
	var stories Story
	//Open file
	jsonFile, err := os.Open("gopher.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	j, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	//Parse JSON file
	if err := json.Unmarshal(j, &stories); err != nil {
		return nil, err
	}
	return stories, nil
}
