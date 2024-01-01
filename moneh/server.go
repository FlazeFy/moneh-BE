package main

import (
	"moneh/packages/database"
	"moneh/routes"
)

func main() {
	database.Init()
	e := routes.InitV1()

	e.Logger.Fatal(e.Start(":1323"))
}
