package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/Michela-DC/book-club/internal/infrastructure/db"
	"github.com/Michela-DC/book-club/internal/infrastructure/webservice"
	"github.com/Michela-DC/book-club/internal/interfaces/controller"
	"github.com/Michela-DC/book-club/internal/usecase/interactor"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}))

	// TODO: read db path from config
	repo, err := db.NewSQLiteBookRepository("database/data/books.db", logger)
	if err != nil {
		panic(err)
	}

	err = repo.ApplyMigrations(context.Background(), "database/migrations")
	if err != nil {
		panic(err)
	}

	i := interactor.NewBookInteractor(repo, logger)
	ctl := controller.NewBookController(i, logger)
	h := webservice.NewHandler(ctl)

	s := &http.Server{
		Addr:              ":8080",
		Handler:           h,
		ReadTimeout:       time.Second,
		WriteTimeout:      2 * time.Second,
		IdleTimeout:       120 * time.Second,
		ReadHeaderTimeout: time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
