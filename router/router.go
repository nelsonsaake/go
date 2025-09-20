package router

import (
	"github.com/gofiber/fiber/v2"
)

// Router is your custom router wrapper
type Router struct {
	group      fiber.Router
	middleware []fiber.Handler
	lastRoute  fiber.Router
}

// NewRouter wraps a fiber.Router
func NewRouter(group fiber.Router) *Router {
	return &Router{group: group}
}

// Use creates a new Router with extra middleware
func (r *Router) Use(handlers ...fiber.Handler) *Router {
	newMiddleware := append([]fiber.Handler{}, r.middleware...) // copy existing
	newMiddleware = append(newMiddleware, handlers...)          // add new ones
	return &Router{
		group:      r.group,
		middleware: newMiddleware,
	}
}

// Add general method to add any HTTP method
func (r *Router) Add(method, path string, handlers ...fiber.Handler) *Router {
	allHandlers := append([]fiber.Handler{}, r.middleware...)
	allHandlers = append(allHandlers, handlers...)

	route := r.group.Add(method, path, allHandlers...)
	return &Router{
		group:      r.group,
		middleware: r.middleware,
		lastRoute:  route,
	}
}

// Post shortcut
func (r *Router) Post(path string, handlers ...fiber.Handler) *Router {
	return r.Add(fiber.MethodPost, path, handlers...)
}

// Put shortcut
func (r *Router) Put(path string, handlers ...fiber.Handler) *Router {
	return r.Add(fiber.MethodPut, path, handlers...)
}

// Get shortcut
func (r *Router) Get(path string, handlers ...fiber.Handler) *Router {
	return r.Add(fiber.MethodGet, path, handlers...)
}

// Delete shortcut
func (r *Router) Delete(path string, handlers ...fiber.Handler) *Router {
	return r.Add(fiber.MethodDelete, path, handlers...)
}

// Patch shortcut
func (r *Router) Patch(path string, handlers ...fiber.Handler) *Router {
	return r.Add(fiber.MethodPatch, path, handlers...)
}

// Options shortcut
func (r *Router) Options(path string, handlers ...fiber.Handler) *Router {
	return r.Add(fiber.MethodOptions, path, handlers...)
}

// Head shortcut
func (r *Router) Head(path string, handlers ...fiber.Handler) *Router {
	return r.Add(fiber.MethodHead, path, handlers...)
}

// Name sets the name on the last added route
func (r *Router) Name(name string) *Router {
	if r.lastRoute != nil {
		r.lastRoute.Name(name)
	}
	return r
}
