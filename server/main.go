package main

import (
	"fmt"
	"github.com/kennyomg/in12tochten/server/routes"
	"github.com/labstack/echo"
)

const port string = ":3333"

func main() {
	fmt.Printf("Running at %v\n", port)

	e := echo.New()

	routes.Init(e)

	e.StartTLS(port, "server/cert/server.crt", "server/cert/server.key")
}
