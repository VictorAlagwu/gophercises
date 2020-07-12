package story

import (
	"encoding/json"
	"fmt"
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

//JSONHandler :
func JSONHandler() (Story, error) {
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
