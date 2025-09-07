package servers

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

type NgrokServer struct {
	App *fiber.App
}

func (s *NgrokServer) Run() error {
	tun, err := ngrok.Listen(context.Background(), config.HTTPEndpoint(), ngrok.WithAuthtokenFromEnv())
	if err != nil {
		return fmt.Errorf("error opening an ngrok tunnel %v", err)
	}
	return s.App.Listener(tun)
}

func NewNgrokServer() *NgrokServer {
	app := fiber.New()
	return &NgrokServer{App: app}
}
