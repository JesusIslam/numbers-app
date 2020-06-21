package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/faldyfin/numbers-app/api"
)

func main() {
	// Echo instance
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
          AllowCredentials: true,
  	  //AllowOrigins: []string{"http://localhost:3000"},
          AllowOrigins: []string{"*"},
  	  AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Set up basic auth with username and password
	e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
		Validator: func(username, password string, c echo.Context) (bool, error) {
			if username == "kita" && password == "bisa" {
				return true, nil
			}
			return false, nil
		},
	}))

	// Route => handler
	e.POST("/numbersapi/:number", api.NumbersHandler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
