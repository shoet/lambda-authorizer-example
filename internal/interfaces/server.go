package interfaces

import "github.com/labstack/echo/v4"

func NewServer() (*echo.Echo, error) {
	e := echo.New()

	e.GET("/health", NewHealthHandler().Handler)

	return e, nil
}

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Handler(c echo.Context) error {
	response := struct {
		Message string `json:"message"`
	}{
		Message: "OK",
	}
	return c.JSON(200, response)
}
