package main

import (
	"encoding/json"
	"fmt"
	"html/template"

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

//NewHandler :
func NewHandler(s Story) http.Handler {
	return handler{s}
}
type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("chapter.html")
	fmt.Println(h.s)
	t.Execute(w, h.s["intro"])
}

func main() {
	//
	story, err := createJSONHandler()
	if err != nil {
		panic(err)
	}
	
	//Start HTTP Server
	fmt.Println("Server Started on Port 8082")
	h := NewHandler(story)
	log.Fatal(http.ListenAndServe(":8082", h))
}

func createJSONHandler() (Story, error) {
	var story Story
	//Open file
	jsonFile, err := os.Open("gopher.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	j := json.NewDecoder(jsonFile)

	//Parse JSON file
	if err := j.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}
