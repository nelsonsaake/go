package servers

import (
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nelsonsaake/go/routes"
)

func (s *FiberServer) Start() error {

	s.App.Use(logger.New())

	m := newManager()

	// add routes to server
	routes.Load(s.App)

	// start server
	errs := make(chan error)

	// start fiber server
	m.Add(s.startFiberServer(errs))

	// start ngrok tunnel
	m.Add(s.exposeViaNgrok(errs))

	// Wait for errors from both servers
	for err := range m.errs {
		if err != nil {
			return err
		}
	}

	return nil
}
