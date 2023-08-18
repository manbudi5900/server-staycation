package main

import (
	routes2 "staycation/routes"
)

func main() {

	routes := routes2.Init()

	routes.Run(":9100")
	// http.ListenAndServe(":1312", routes)
}