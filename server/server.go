package server

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tumbleweedd/avito-test-task/pkg/handler"
	"github.com/tumbleweedd/avito-test-task/pkg/repository"
	"github.com/tumbleweedd/avito-test-task/pkg/service"
)

type Server struct {
	handler    *handler.Handler
	repository *repository.Repository
	service    *service.Service
	echo       *echo.Echo
}

func NewServer(db *sqlx.DB) (*Server, error) {
	server := &Server{}

	server.repository = repository.NewRepository(db)
	server.service = service.NewService(server.repository)
	server.handler = handler.NewHandler(server.service)

	server.echo = echo.New()

	server.handler.InitRoutes(server.echo)

	return server, nil

}

func (a *Server) Run() error {
	fmt.Println("server running")

	err := a.echo.Start(":8000")
	if err != nil {
		logrus.Fatal(err)
	}

	return nil
}
