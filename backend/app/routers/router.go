package router

import (
	controller "app/controllers"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {

	var allowedOrigins = []string{
		"http://127.0.0.1:3000",
		"http://localhost:3000",
	}

	var allowHeaders = []string{
		echo.HeaderAccessControlAllowHeaders,
		echo.HeaderContentType,
		echo.HeaderContentLength,
		echo.HeaderAcceptEncoding,
		echo.HeaderXCSRFToken,
		echo.HeaderAuthorization,
	}

	var allowMethods = []string{
		http.MethodGet,
		http.MethodPut,
		http.MethodPatch,
		http.MethodPost,
		http.MethodDelete,
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     allowedOrigins,
		AllowHeaders:     allowHeaders,
		AllowMethods:     allowMethods,
		MaxAge:           86400,
	}))

	// req,resの値確認
	e.Use(middleware.BodyDump(func(_ echo.Context, reqBody, resBody []byte) {
		fmt.Fprintf(os.Stderr, "Request: %v\n", string(reqBody))
		fmt.Fprintf(os.Stderr, "Response: %v\n", string(resBody))
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("todos", controller.TodoRead())
	e.POST("todos/create", controller.TodoCreate())
	e.PUT("todos/:id", controller.TodoUpdate())
	e.DELETE("todos/:id", controller.TodoDelete())

	return e
}
