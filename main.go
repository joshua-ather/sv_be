package main

import "github.com/joshua-ather/sv_be/routes"

func main() {
	route := routes.Init()

	route.Logger.Fatal(route.Start(":7444"))
}
