package interfaces

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shoet/lambda-authorizer-example/internal/infrastracture"
	"github.com/shoet/lambda-authorizer-example/internal/service"
)

func NewServer() (*echo.Echo, error) {
	e := echo.New()

	kvs := infrastracture.NewKeyValueStore()
	authService := service.NewAuthService(kvs)

	e.GET("/health", NewHealthHandler().Handler)
	e.POST("/login", NewLoginHandler(authService).Handler)

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

type AuthService interface {
	Login(ctx context.Context, email, password string) (bool, error)
	GenerateToken(ctx context.Context) (string, error)
}

type LoginHandler struct {
	AuthService AuthService
}

func NewLoginHandler(authService AuthService) *LoginHandler {
	return &LoginHandler{AuthService: authService}
}

func (h *LoginHandler) Handler(c echo.Context) error {
	var requestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	defer c.Request().Body.Close()
	if err := json.NewDecoder(c.Request().Body).Decode(&requestBody); err != nil {
		return c.JSON(400, ErrorResponse{Message: http.StatusText(400)})
	}
	if requestBody.Email == "" || requestBody.Password == "" {
		return c.JSON(400, ErrorResponse{Message: http.StatusText(400)})
	}
	ok, err := h.AuthService.Login(c.Request().Context(), requestBody.Email, requestBody.Password)
	if err != nil {
		return c.JSON(500, ErrorResponse{Message: http.StatusText(500)})
	}
	if !ok {
		return c.JSON(401, ErrorResponse{Message: http.StatusText(401)})
	}
	token, err := h.AuthService.GenerateToken(c.Request().Context())
	if err != nil {
		return c.JSON(500, ErrorResponse{Message: http.StatusText(500)})
	}
	return c.JSON(200, struct {
		AuthToken string `json:"auth_token"`
	}{
		AuthToken: token,
	})
}

type ErrorResponse struct {
	Message string `json:"message"`
}
