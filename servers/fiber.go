package servers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type FiberServer struct {
	App  *fiber.App
	Port string
}

func (s *FiberServer) Run() error {
	return s.App.Listen(cleanPort(s.Port))
}

func (s *FiberServer) NewNgrokServer() Server {
	return &NgrokServer{App: s.App}
}

func NewFiberServer(port string) *FiberServer {
	app := fiber.New()
	app.Use(cors.New())
	return &FiberServer{App: app, Port: port}
}
