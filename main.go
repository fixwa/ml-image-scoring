package main

import (
	"test/tagger"
	"test/webserver"
)

func main() {
	t := tagger.NewTagger(
		"./scoring_class_index.json",
		"http://scoring_api:5001/score",
	)

	server := webserver.NewServer(
		"9999",
		5120*1024, // 5MB
		t,
	)

	server.Serve()
}
