package main

import (
	"net/http"
	routes2 "staycation/routes"
)

func main() {

	routes := routes2.Init()

	// routes.Run(":8000")
	http.ListenAndServe(":8000", routes)
}