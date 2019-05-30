package main

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"godejs/api/ws"
	"net/http"
	"godejs/api/calculateLag"
)

var (
	upgrader = websocket.Upgrader{}
)

func main() {
	calculateLag.Test()
	e := echo.New()
	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Use(middleware.CORS())
	e.GET("/ws", ws.Listen)
	e.File("/", "build/index.html")
	e.Static("/static", "build/static")
	e.Logger.Fatal(e.Start(":1323"))
}
