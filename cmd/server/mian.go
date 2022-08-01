package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	svc "news-app/internal/service"
	"news-app/internal/store"

	"news-app/internal/cache"
	"news-app/internal/config"
	"news-app/internal/news"
	"news-app/internal/transport/http"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatal(err)
	}

	db := createDatabaseConnection(cfg.DB)
	st := store.New(db)

	service := svc.New(st)

	fetcher := news.NewFetcher(cfg.NewsSources)
	cachedFetcher := cache.NewCachedNewsFetcher(cfg.NewsSources, fetcher, cfg.CacheTTL)

	handler := http.NewHandler(cachedFetcher, cfg.Port, service)

	handler.Serve()
}

func createDatabaseConnection(cfg *config.Database) *sqlx.DB {
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName,
	)
	DB, err := sqlx.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	return DB
}
