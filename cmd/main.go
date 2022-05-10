package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ammiranda/xkcdscraper/models"
	"github.com/ammiranda/xkcdscraper/repository"
)

const (
	xkcd_base_url = "https://xkcd.com/%s/info.0.json"
)

func main() {
	repo, err := repository.NewXkcdComicRepository()
	if err != nil {
		panic(err)
	}

	comicCount := 1

	for {
		resp, err := http.Get(fmt.Sprintf("https://xkcd.com/%d/info.0.json", comicCount))
		if err != nil {
			panic(err)
		}

		var comic models.XkcdComic

		json.NewDecoder(resp.Body).Decode(&comic)

		err = repo.AddComic(&comic)
		if err != nil {
			log.Print(fmt.Sprintf("Comic %d failed to save", comic.Num))
		}
		comicCount += 1
	}
}
