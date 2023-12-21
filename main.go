package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ikun666/gikun/gikun"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello,ikun")
}
func Test(w http.ResponseWriter, req *http.Request) {
	time.Sleep(100 * time.Millisecond)
	fmt.Fprintf(w, "test")
}
func main() {
	r := gikun.New()
	r.GET("/hello", Hello)
	r.POST("/test", Test)
	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
