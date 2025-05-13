package server

import (
	"backend/config"
	"backend/logger"
	"backend/models"
	"backend/repository"
	"encoding/json"
	"fmt"
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
	PeopleRepository repository.Repository[models.Person]
	KillRepository   repository.Repository[models.Kill]
	logger           *logger.Logger
	taskQueue        *TaskQueue
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
		Handler: s.router(),
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
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
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
