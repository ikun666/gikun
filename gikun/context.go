package gikun

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]any

// gikun context
type Context struct {
	//原始http参数
	Writer http.ResponseWriter
	Req    *http.Request

	//请求参数
	Path   string
	Method string
	Params map[string]string //解析后的参数

	//响应状态
	StatusCode int

	//中间件
	handler []HandlerFunc
	index   int //记录运行到第几个handler
}

// 创建Context实例
func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index:  -1,
	}
}

// Next() 转到下一个中间件
func (c *Context) Next() {
	//确保每次都会++
	c.index++
	l := len(c.handler)
	for ; c.index < l; c.index++ {
		//执行handler没有Next()也能++
		c.handler[c.index](c)
	}
}

// 获取解析路由参数
func (c *Context) Param(key string) string {
	value := c.Params[key]
	return value
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
