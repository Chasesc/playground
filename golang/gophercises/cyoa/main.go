package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	DEFAULT_STORY_JSON_FILENAME    = "gopher.json"
	DEFAULT_HTML_TEMPLATE_FILENAME = "layout.html"
	HTTP_PORT                      = 8080
	KEY_STORY_BEGINS_AT            = "intro"
)

type storyBlock struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func loadStoryJson(filename string) map[string]storyBlock {
	fileData, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var story map[string]storyBlock
	if err := json.Unmarshal(fileData, &story); err != nil {
		log.Fatal(err)
	}

	return story
}

func loadStoryHTMLTemplate(filename string) *template.Template {
	return template.Must(template.ParseFiles(filename))
}

func getStoryBlockKeyFromRequestURI(uri string) string {
	storyBlockKey := uri[1:]
	if uri == "/" {
		storyBlockKey = KEY_STORY_BEGINS_AT
	}
	return storyBlockKey
}

func createStoryHTTPRequestHandler(tmpl *template.Template, story map[string]storyBlock) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		storyBlockKey := getStoryBlockKeyFromRequestURI(req.RequestURI)

		if requestedStoryBlock, ok := story[storyBlockKey]; ok {
			tmpl.Execute(w, requestedStoryBlock)
			return
		}

		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Unknown story block!")
	}
}

func main() {
	storyFilenamePtr := flag.String("storyFilename", DEFAULT_STORY_JSON_FILENAME, "Story JSON filename.")
	templateFilenamePtr := flag.String("templateFilename", DEFAULT_HTML_TEMPLATE_FILENAME, "HTML Story Template.")
	flag.Parse()

	storyMapping := loadStoryJson(*storyFilenamePtr)
	tmpl := loadStoryHTMLTemplate(*templateFilenamePtr)

	handler := createStoryHTTPRequestHandler(tmpl, storyMapping)

	addr := fmt.Sprintf(":%d", HTTP_PORT)
	fmt.Printf("Starting the server on %s\n", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal(err)
	}
}
