package servers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nelsonsaake/go/routes"
)

func (s *FiberServer) startFiberServer(errs chan error) {

	go func() {
		errs <- s.Run()
	}()
}

func (s *FiberServer) exposeViaNgrok(errs chan error) {

	if s.UseNgrok != "true" {
		log.Println("Ngrok tunnel not enabled. Skipping ngrok setup.")
		return
	}

	log.Println("Ngrok tunnel enabled. Setting up ngrok...")
	ngrok := s.NewNgrokServer()

	go func() {
		errs <- ngrok.Run()
	}()
}

func (s *FiberServer) Start() error {

	log.Println("Starting Fiber server...")
	fmt.Println("Starting Fiber server...")
	fmt.Printf("Server will listen on port: %s\n", s.Port)

	s.App.Use(logger.New())

	// add routes to server
	routes.Load(s.App)

	// start server
	errs := make(chan error)

	// start fiber server
	s.startFiberServer(errs)

	// start ngrok tunnel
	s.exposeViaNgrok(errs)

	// Wait for errors from both servers
	for err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}
