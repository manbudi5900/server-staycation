package main

import (
	routes2 "staycation/routes"
)

func main() {

	routes := routes2.Init()

	routes.Run()
}