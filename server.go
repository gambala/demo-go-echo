package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Root endpoint
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// JSON endpoint
	e.GET("/json", func(c echo.Context) error {
		response := map[string]string{"message": "Hello, JSON!"}
		return c.JSON(http.StatusOK, response)
	})

	// HTML endpoint
	e.GET("/html", func(c echo.Context) error {
		htmlContent := "<html><body><h1>Hello, HTML!</h1></body></html>"
		return c.HTML(http.StatusOK, htmlContent)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
