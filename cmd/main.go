package main

import (
	"fmt"

	"github.com/anomaly44/go-tmpl/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	userHandler := handler.UserHandler{}
	app.Static("/static", "static")
	app.GET("/", userHandler.HandleLoginShow)
	app.POST("/login", userHandler.HandleLoginSubmit)
	app.GET("/user", userHandler.HandleUserShow)
	fmt.Println("It is working")

	app.Start(":8080")
}
