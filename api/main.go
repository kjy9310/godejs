package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"strconv"
	"time"
)

var (
	clients = make(map[*websocket.Conn]bool)
	upgrader = websocket.Upgrader{}
	channel = make(chan string)
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
		channel <- string(msg)
		fmt.Printf("%s\n", msg)
		unixStampString := msg[0:10]
		fmt.Println("unixStampString : ",string(unixStampString))
		unixStamp, err := strconv.ParseInt(string(unixStampString), 10, 64)
		if (err != nil){
			fmt.Println("!!!!ERR!!!!", err)
		}
		fmt.Println(time.Unix(int64(unixStamp), 0))
	}
}

func writeAll(){
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
