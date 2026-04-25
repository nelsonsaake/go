package servers

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

type NgrokServer struct {
	App *fiber.App
}

func (s *NgrokServer) Run() error {

	var (
		ctx               = context.Background()
		ngrokConfig       = config.HTTPEndpoint()
		ngrokTokenFromEnv = ngrok.WithAuthtokenFromEnv()
	)

	tun, err := ngrok.Listen(ctx, ngrokConfig, ngrokTokenFromEnv)
	if err != nil {
		return fmt.Errorf("error opening an ngrok tunnel %v", err)
	}

	log.Println("Exposing server via ngrok...")
	log.Println(ngrokConfig)
	log.Println(ngrokTokenFromEnv)

	return s.App.Listener(tun)
}

func NewNgrokServer() *NgrokServer {
	app := fiber.New()
	return &NgrokServer{App: app}
}
