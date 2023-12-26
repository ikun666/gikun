package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ikun666/gikun/gikun"
)

func Hello(c *gikun.Context) {
	c.SendString(200, fmt.Sprintf("hello,%s path:%s", c.Param("name"), c.Path))
}
func JSON(c *gikun.Context) {
	c.SendJSON(201, gikun.H{
		"ikun":   666,
		"go_web": "gikun",
	})
}
func Data(c *gikun.Context) {
	c.SendData(202, []byte(fmt.Sprintf("data:%s", c.Param("data"))))
}
func HTML(c *gikun.Context) {
	c.SendHTML(203, "<!DOCTYPE html><html><head><meta charset='utf-8'><title>ZONGXP</title></head><body><p>gikun 家人们，太强了</p></body></html>")
}
func Logger(c *gikun.Context) {
	t := time.Now()
	c.Next()
	log.Printf("[%d] %s in %v\n", c.StatusCode, c.Req.RequestURI, time.Since(t))
}
func TestFunc(c *gikun.Context) {
	log.Printf("%s\n", c.Req.RequestURI)
}
func main() {
	r := gikun.New()
	v1 := r.Group("/v1").Use(Logger)
	{
		v1.GET("/hello/:name", Hello)
		v1.POST("/html", HTML)
	}
	v2 := r.Group("/v2").Use(TestFunc)
	{
		v2.GET("/json", JSON)
		v2.POST("/data/*data", Data)
	}

	v3 := v2.Group("/v3")
	{
		v3.GET("/hello/:name", Hello)
		v3.POST("/data/*data", Data)
	}

	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
