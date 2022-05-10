package repository

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ammiranda/xkcdscraper/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type XkcdComicRepository struct {
	db *gorm.DB
}

func NewXkcdComicRepository() (*XkcdComicRepository, error) {
	host := os.Getenv("HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	sslMode := os.Getenv("POSTGRES_SSL_MODE")
	timezone := os.Getenv("POSTGRES_TIMEZONE")
	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s Timezone=%s", host, user, password, dbname, port, sslMode, timezone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &XkcdComicRepository{
		db: db,
	}, nil
}

func (x *XkcdComicRepository) AddComic(comic *models.XkcdComic) error {
	return x.db.Create(comic).Error
}
