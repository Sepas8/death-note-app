package main

import (
	"log"
	"github.com/Sepas8/death-note-app/backend/config"
	"github.com/Sepas8/death-note-app/backend/logger"
	"github.com/Sepas8/death-note-app/backend/repository"
	"github.com/Sepas8/death-note-app/backend/server"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Cargar configuración
	cfg := config.Load()

	// Inicializar logger
	logger.Init(cfg.Environment)

	// Conectar a PostgreSQL
	dsn := repository.BuildDSN(cfg.DB)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrar modelos
	if err := repository.Migrate(db); err != nil {
		log.Fatalf("Failed to migrate models: %v", err)
	}

	// Inicializar servidor
	srv := server.NewServer(cfg, db)

	// Iniciar el servidor
	log.Printf("🚀 Server running on %s", cfg.ServerAddress)
	if err := srv.Start(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}