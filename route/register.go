package Route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

type Handler fiber.Handler

type Router = fiber.Router

type Item struct {
	Method   string
	Path     string
	Handlers []fiber.Handler
}

var registry = []Item{}

func loadRegister(c *cobra.Command) {
	// for name, handler := range registry {
	// 	// register(c, name, handler)
	// }
}

// add general method to add any HTTP method
func add(method, path string, handlers ...fiber.Handler) {
	registry = append(registry, Item{
		Method:   method,
		Path:     path,
		Handlers: handlers,
	})
}

// Post
func Post(path string, handlers ...fiber.Handler) {
	add(fiber.MethodPost, path, handlers...)
}

// Put
func Put(path string, handlers ...fiber.Handler) {
	add(fiber.MethodPut, path, handlers...)
}

// Get
func Get(path string, handlers ...fiber.Handler) {
	add(fiber.MethodGet, path, handlers...)
}

// Delete
func Delete(path string, handlers ...fiber.Handler) {
	add(fiber.MethodDelete, path, handlers...)
}

// Patch
func Patch(path string, handlers ...fiber.Handler) {
	add(fiber.MethodPatch, path, handlers...)
}

// Options
func Options(path string, handlers ...fiber.Handler) {
	add(fiber.MethodOptions, path, handlers...)
}

// Head
func Head(path string, handlers ...fiber.Handler) {
	add(fiber.MethodHead, path, handlers...)
}
