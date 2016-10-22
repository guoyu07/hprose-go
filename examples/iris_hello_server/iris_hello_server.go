package main

import (
	"github.com/hprose/hprose-golang/rpc"
	"github.com/kataras/iris"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	service := rpc.NewFastHTTPService()
	service.AddFunction("hello", hello)
	iris.Any("/hello", func(c *iris.Context) {
		service.ServeFastHTTP(c.RequestCtx)
	})
	iris.Listen(":8080")
}
