package main

import (
	"flag"
	"sync"

	"github.com/khatibomar/paytabs-challenge/internal/store"
	"go.uber.org/zap"
)

type Config struct {
	port int
	env  string
}

type application struct {
	config Config
	logger *zap.Logger
	store  store.Store
	wg     sync.WaitGroup
}

const version = "1.0.0"

func main() {
	var cfg Config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	var logger *zap.Logger
	if cfg.env == "development" {
		logger = zap.Must(zap.NewDevelopment())
	} else {
		logger = zap.Must(zap.NewProduction())
	}
	defer logger.Sync()

	s := store.New()
	logger.Info("Ingesting accounts into the store...")
	err := s.Seed("./data/accounts-mock.json")
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("Ingesting accounts done...")
	app := &application{
		config: cfg,
		logger: logger,
		store:  s,
	}

	if err := app.serve(); err != nil {
		app.logger.Fatal(err.Error())
	}
}
