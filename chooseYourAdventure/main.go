package main

import (
	"fmt"
	"html/template"
	"strings"

	"log"
	"net/http"

	"github.com/victoralagwu/gophercises/chooseYourAdventure/story"
)

//NewHandler :
func NewHandler(s story.Story) http.Handler {
	return handler{s}
}
type handler struct {
	s story.Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("chapter.html")
	currentPath := strings.TrimSpace(r.URL.Path)

	if currentPath == "" || currentPath == "/" {
		currentPath = "intro"
	} else {
		currentPath = currentPath[1:]
	}
	
	if chapter, ok := h.s[currentPath]; ok {
		err := t.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found.", http.StatusNotFound)
} 

func main() {
	story, err := story.JSONHandler()
	if err != nil {
		panic(err)
	}
	
	fmt.Println("Server Started on Port 8082")
	h := NewHandler(story)
	log.Fatal(http.ListenAndServe(":8082", h))
}
