package main

import (
	"fmt"

	"aries-technical-challenge/routes"
)

func main() {
	ginRouter := routes.SetupRouter()
	fmt.Println("Starting server on port 8080")
	ginRouter.Run(":8080")
}
