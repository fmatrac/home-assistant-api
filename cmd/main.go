package main

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"

	c "github.com/fmatrac/home-assistant-api/config"
	"github.com/fmatrac/home-assistant-api/pkg/api"
	a "github.com/fmatrac/home-assistant-api/pkg/application"
	l "github.com/fmatrac/home-assistant-api/pkg/logger"
	"github.com/fmatrac/home-assistant-api/pkg/postgres"
	"github.com/fmatrac/home-assistant-api/web"
	_ "github.com/lib/pq"
)

func run(ctx context.Context, w io.Writer, args []string) error {
	cfg, err := c.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}
	logger := l.Setup(cfg)

	// Connect to database
	db, err := sql.Open("postgres", cfg.Postgres.Conn)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	// Test connection
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}
	logger.Info("Connected to database")

	// Initialize repositories
	wydarzeniaRepo := postgres.NewWydarzeniaKalendarzRepository(db)
	przypomnieniRepo := postgres.NewPrzypomnieniRepository(db)
	produktyRepo := postgres.NewProduktyRepository(db)
	listyZakupowRepo := postgres.NewListyZakupowRepository(db)
	pozycjeListyZakupowRepo := postgres.NewPozycjeListyZakupowRepository(db)
	stanyMagazynoweRepo := postgres.NewStanyMagazynoweRepository(db)
	historiaStanuZapasowRepo := postgres.NewHistoriaStanuZapasowRepository(db)

	// Create application
	application := a.New(
		cfg,
		logger,
		wydarzeniaRepo,
		przypomnieniRepo,
		produktyRepo,
		listyZakupowRepo,
		pozycjeListyZakupowRepo,
		stanyMagazynoweRepo,
		historiaStanuZapasowRepo,
	)

	// Initialize server with embedded frontend
	api.InitServer(cfg, logger, application, http.FS(web.FS))

	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
