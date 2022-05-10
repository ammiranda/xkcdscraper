package models

import (
	"os"

	"gorm.io/gorm"
)

type XkcdComic struct {
	gorm.Model
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func (c *XkcdComic) TableName() string {
	tableName := os.Getenv("XKCD_COMIC_TABLE_NAME")
	return tableName
}
