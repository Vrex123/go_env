package main

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
	"net/http"
)

type Config struct {
	Addr   string `env:"ADDR" env-default:":3000"`
	DbHost string `env:"DB_HOST"`
	DbPort string `env:"DB_PORT" env-default:"5432"`
}

func main() {
	var cfg Config

	err := cleanenv.ReadEnv(&cfg)

	if err != nil {
		slog.Error("Failed to read config", "err", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	slog.Info("Starting server", "addr", cfg.Addr)
	slog.Info("Database params", "host", cfg.DbHost, "port", cfg.DbPort)
	err = http.ListenAndServe(cfg.Addr, mux)

	if err != nil {
		slog.Error("Failed to start server", "err", err)
	}
}
