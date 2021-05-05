package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/ryechus/cyoa"
)

func main() {
	var filename string
	flag.StringVar(&filename, "filename", "story.json", "File that contains all of the chapters of the story")
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	story := cyoa.JsonStory(file)

	http.ListenAndServe("127.0.0.1:3000", cyoa.StoryHandler(story))
}
