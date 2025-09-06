package server

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/nelsonsaake/go/prt"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

type Server interface {
	Run() error
}

type server struct {
	App  *fiber.App
	Port string
}

func (s *server) Run() error {
	return s.App.Listen(prt.Clean(s.Port))
}

func (s *server) Ngrok() error {
	tun, err := ngrok.Listen(
		context.Background(),
		config.HTTPEndpoint(),
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		return fmt.Errorf("error opening an ngrok tunnel%w", err)
	}
	return s.App.Listener(tun)
}

func (s *server) WithPort(port string) *server {
	s.Port = port
	return s
}

func New(port ...string) *server {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	p := ""
	if len(port) > 0 {
		p = port[0]
	}
	return &server{App: app, Port: p}
}
