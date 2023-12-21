package gikun

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]any

// gikun context
type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	StatusCode int
}

// 创建Context实例
func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// 获取POST参数
func (c *Context) GetPOSTValue(key string) string {
	return c.Req.FormValue(key)
}

// 获取GET参数
func (c *Context) GetGETValue(key string) string {
	return c.Req.URL.Query().Get(key)
}

// 设置返回状态码
func (c *Context) SetStatus(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// 设置消息头
func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

// string响应
func (c *Context) SendString(code int, str string) {
	c.SetHeader("Content-Type", "text/plain")
	c.SetStatus(code)
	c.Writer.Write([]byte(str))
}

// json响应
func (c *Context) SendJSON(code int, obj any) {
	c.SetHeader("Content-Type", "application/json")
	c.SetStatus(code)
	data, err := json.Marshal(obj)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Writer.Write([]byte(data))
}

// data响应
func (c *Context) SendData(code int, data []byte) {
	c.SetStatus(code)
	c.Writer.Write(data)
}

// html响应
func (c *Context) SendHTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.SetStatus(code)
	c.Writer.Write([]byte(html))
}
