package main

import (
	"fmt"

	"github.com/anomaly44/go-tmpl/handler"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// In production, the secret key of the CookieStore
// and database name would be obtained from a .env file
const (
	SECRET_KEY string = "secret"
	DB_NAME    string = "app_data.db"
)

func main() {
	app := echo.New()

	userHandler := handler.UserHandler{}

	// logger middleware
	app.Use(middleware.Logger())

	// Session middleware
	app.Use(session.Middleware(sessions.NewCookieStore([]byte(SECRET_KEY))))

	app.Static("/static", "static")
	app.GET("/", userHandler.HandleLoginShow)
	app.POST("/login", userHandler.HandleLoginSubmit)

	proctectedGroup := app.Group("/users", userHandler.AuthMiddleware)
	proctectedGroup.GET("", userHandler.HandleUserShow)
	fmt.Println("It is working")

	app.Start(":8080")
}
