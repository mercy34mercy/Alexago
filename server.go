package main

import (
	"bytes"
	"net/url"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
)


func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		userid := c.QueryParam("userid")
		text  :=  c.QueryParam("prompt")
		print(text)
		go fetch(userid,text)
		time.Sleep(3000)
		return c.String(http.StatusOK, "OK")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func fetch(userid string,text string)(string){
	print(text,userid)
	url := "http://34.82.102.69:5000?userid=" + userid +"&prompt=" + url.QueryEscape(text)
	req, _ := http.NewRequest("GET", url, nil)
	client := new(http.Client)
	r, _ := client.Do(req)
	print(r)
	return "OK"
}




