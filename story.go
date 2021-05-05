package cyoa

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"text/template"
)

type Chapter struct {
	Title      string   `json:"title,omitempty"`
	Paragraphs []string `json:"paragraphs,omitempty"`
	Options    []Option
}

type Option struct {
	Text    string `json:"text,omitempty"`
	Chapter string `json:"chapter,omitempty"`
}

type Story map[string]Chapter

var tpl *template.Template

func init() {
	tpl, _ = template.ParseFiles("templates/cyoa.html")
	// if err != nil {
	// 	panic(err)
	// }
}

func JsonStory(r io.Reader) Story {
	var story Story
	d := json.NewDecoder(r)
	if err := d.Decode(&story); err != nil {
		panic(err)
	}
	return story
}

func StoryHandler(s Story) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url_path := strings.TrimSpace(r.URL.Path)[1:]
		if chapter, ok := s[url_path]; ok {
			tpl.Execute(w, chapter)
			return
		}
		w.Write([]byte("None"))
		return
	}
}
