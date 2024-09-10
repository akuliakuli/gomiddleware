package app

import (
	"fmt"
	"log"

	"github.com/akuliakuli/gomiddleware/internal/app/endpoint"
	"github.com/akuliakuli/gomiddleware/internal/app/mw"
	"github.com/akuliakuli/gomiddleware/internal/app/service"
	"github.com/labstack/echo/v4"
)

type App struct{
	e *endpoint.Endpoint
	echo *echo.Echo
	s *service.Service
}

func New() (*App, error){
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	fmt.Println("Server running")

	a.echo = echo.New()

	a.echo.Use(mw.RoleCheck)

	a.echo.GET("/status",a.e.Status)
	
	return a, nil
	
}

func (a *App) Run() error{
	fmt.Println("Server running")

	err := a.echo.Start(":8000")
	if err != nil{
		log.Fatal(err)
	}

	return nil
}