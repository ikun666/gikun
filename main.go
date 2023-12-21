package main

import (
	"fmt"

	"github.com/ikun666/gikun/gikun"
)

func String(c *gikun.Context) {
	c.SendString(200, "string")
}
func JSON(c *gikun.Context) {
	c.SendJSON(201, gikun.H{
		"ikun":   666,
		"go_web": "gikun",
	})
}
func Data(c *gikun.Context) {
	c.SendData(202, []byte("data~"))
}
func HTML(c *gikun.Context) {
	c.SendHTML(203, "<!DOCTYPE html><html><head><meta charset='utf-8'><title>ZONGXP</title></head><body><p>gikun 家人们，太强了</p></body></html>")
}
func main() {
	r := gikun.New()
	r.GET("/string", String)
	r.GET("/json", JSON)
	r.POST("/data", Data)
	r.POST("/html", HTML)
	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
