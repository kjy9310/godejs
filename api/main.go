package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Use(middleware.CORS())

	e.File("/", "build/index.html")
	e.GET("/faqs", func(c echo.Context) error {
		return c.File("/Users/kimjuyoun/dev/test/webサイト-質問-回答reduced.csv")
	})
	e.GET("/examples", func(c echo.Context) error {
		return c.File("/Users/kimjuyoun/dev/test/インポート-質問例.csv")
	})
	e.GET("/dictionaries", func(c echo.Context) error {
		return c.File("/Users/kimjuyoun/dev/test/インポート-辞書.csv")
	})
	//e.Static("/csv", "/Users/kimjuyoun/dev/test/")
	e.Logger.Fatal(e.Start(":1323"))
}
