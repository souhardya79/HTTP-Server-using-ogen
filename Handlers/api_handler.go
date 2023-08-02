package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type APIHandler struct{}

func NewAPIHandler() *APIHandler {
	return &APIHandler{}
}

func (h *APIHandler) GetHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!")
}
