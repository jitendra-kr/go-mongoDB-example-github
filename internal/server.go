package internal

import (
	"fmt"

	"github.com/go-mongoDB-example-github/db"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	Server *echo.Echo
	Store  *mongo.Database
}

func New() *Server {
	e := echo.New()
	db := db.ConnectDB()
	s := Server{
		Server: e,
		Store:  db,
	}
	s.AddRoutes()
	return &s
}

func (w *Server) Start() error {
	if err := w.Server.Start(":9090"); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
