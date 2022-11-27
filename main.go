package main

import (
	"fmt"
	"net/http"

	"github.com/Rudolphlin/crw/engine"
)

func main() {
	e := engine.New()
	e.Get("/", indexHandler)
	e.Get("/hello", helloHandler)
	e.Get("/helloJson", helloJsonHandler)
	e.Post("/hello", helloHandler)

	fmt.Print(e.Run(":8080"))
}

func indexHandler(c *engine.Context) {
	c.HTML(http.StatusOK, "<h1>Welcome to crw web</h1>")
}

// handler echoes r.URL.Header
func helloHandler(c *engine.Context) {
	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
}

func helloJsonHandler(c *engine.Context) {
	c.JSON(http.StatusOK, fmt.Sprintf("{\"name\":%s,\"age\":%s}", c.Query("name"), c.Query("age")))
}
