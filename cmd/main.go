package main

import (
	"webframe"
)

func main() {
	engine := webframe.NewEngine()
	g := engine.NewGroup("/hello")
	g.GET("", hello)
	g.GET("/hello", hello)
	engine.GET("/hello2", hello2)
	engine.Run(":8080")
}

func hello(c *webframe.Context) {
	c.StringResponse(200, "Hello")
}

func hello2(c *webframe.Context) {
	c.JSONResponse(200, webframe.M{"hello": "world"})
}
