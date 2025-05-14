package server

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Sepas8/death-note-app/backend/config"
	"github.com/Sepas8/death-note-app/backend/logger"
	"github.com/Sepas8/death-note-app/backend/models"
	"github.com/Sepas8/death-note-app/backend/repository"

	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Server struct {
	DB               *gorm.DB
	Config           *config.Config
	Handler          http.Handler
	PeopleRepository *repository.PeopleRepository
	KillRepository   *repository.KillRepository
	logger           *logger.Logger
	taskQueue        *TaskQueue
}

func NewTaskQueue() *TaskQueue {
	return &TaskQueue{
		tasks: make(map[int]context.CancelFunc),
	}
}

func NewServer() *Server {
	s := &Server{
		logger:    logger.NewLogger(),
		taskQueue: NewTaskQueue(),
	}
	var config config.Config
	configFile, err := os.ReadFile("config/config.json")
	if err != nil {
		s.logger.Fatal(err)
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		s.logger.Fatal(err)
	}
	s.Config = &config
	return s
}

func (s *Server) StartServer() {
	fmt.Println("Inicializando base de datos...")
	s.initDB()
	fmt.Println("Inicializando mux...")
	srv := &http.Server{
		Addr:    s.Config.Address,
		Handler: s.Router(),
	}
	fmt.Println("Escuchando en el puerto ", s.Config.Address)
	if err := srv.ListenAndServe(); err != nil {
		s.logger.Fatal(err)
	}
}

func (s *Server) initDB() {
	switch s.Config.Database {
	case "sqlite":
		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			s.logger.Fatal(err)
		}
		s.DB = db
	case "postgres":
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			s.Config.DBHost,
			s.Config.DBPort,
			s.Config.DBUser,
			s.Config.DBPassword,
			s.Config.DBName,
		)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			s.logger.Fatal(err)
		}
		s.DB = db
	}
	fmt.Println("Aplicando migraciones...")
	s.DB.AutoMigrate(&models.Person{}, &models.Kill{})
	s.KillRepository = repository.NewKillRepository(s.DB)
	s.PeopleRepository = repository.NewPeopleRepository(s.DB)
}
