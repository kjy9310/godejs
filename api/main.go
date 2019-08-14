package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"strconv"
	"time"

	"godejs/api/pingSync"
)

var (
	clients  = make(map[*websocket.Conn]bool)
	upgrader = websocket.Upgrader{}
	channel  = make(chan string)
)

func hello(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()
	clients[ws] = true
	ws.WriteMessage(websocket.TextMessage, []byte("connected"))
	for {
		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
		unixStampString := msg[0:13]
		fmt.Println("client unixStampString : ", string(unixStampString))
		unixStamp, err := strconv.ParseInt(string(unixStampString), 10, 64)
		if err != nil {
			fmt.Println("!!!!ERR!!!!", err)
		}
		fmt.Println("client : ", time.Unix(0, int64(unixStamp)*int64(time.Millisecond)))
		fmt.Println("servertime : ", int64(time.Now().UnixNano()/1000000))
		fmt.Println("server : ", time.Unix(0, int64(time.Now().UnixNano()/1000000)*int64(time.Millisecond)))
		channel <- string(strconv.Itoa(int(time.Now().UnixNano()/1000000)))
	}
}

func writeAll() {
	for {
		for ws := range clients {
			err := ws.WriteMessage(websocket.TextMessage, []byte(<-channel))
			if err != nil {
				fmt.Println(err)
				ws.Close()
				delete(clients, ws)
			}
		}
	}
}

func main() {
	pingSync.Test()

	e := echo.New()

	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Use(middleware.CORS())

	e.File("/", "build/index.html")
	e.Static("/static", "build/static")
	e.GET("/ws", hello)
	go writeAll()
	e.Logger.Fatal(e.Start(":1323"))
}
