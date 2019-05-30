package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"strconv"
	"time"
)
var (
	upgrader = websocket.Upgrader{}
)
func Listen(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}

		fmt.Printf("%s\n", msg)
		unixStampString := msg[0:10]
		fmt.Println("unixStampString : ", string(unixStampString))
		unixStamp, err := strconv.ParseInt(string(unixStampString), 10, 64)
		if err != nil {
			fmt.Println("!!!!ERR!!!!", err)
		}
		fmt.Println(time.Unix(int64(unixStamp), 0))
	}
}
