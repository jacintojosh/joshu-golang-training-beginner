package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type healthCheckResponse struct {
	Status string `json:"status"`
}

type helloWorldResponse struct {
	Message string `json:"message"`
}

type serverHandler struct{}

// Defines routes and initializes handlers
func InitServerHandler(e *echo.Echo) {
	// Controllers for the Handlers entrypoint
	h := &serverHandler{}
	// Structure is route, <handler-function>, <middleware...>
	e.GET("/health", h.Health)
	e.GET("/hello-world", h.HelloWorld)
}

// Handler function for route /health
func (h *serverHandler) Health(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, healthCheckResponse{
		Status: "healthy",
	})
}

// Handler function for route /hello-world
func (h *serverHandler) HelloWorld(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, helloWorldResponse{
		Message: "hello world",
	})
}
