package main

import (
	//"moneh/factories"

	"moneh/packages/database"
	"moneh/packages/telegram"
	"moneh/routes"
	"sync"
)

func main() {
	// Run App
	database.Init()

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		telegram.InitTeleBot()
	}()

	go func() {
		defer wg.Done()
		e := routes.InitV1()
		if err := e.Start(":1324"); err != nil {
			e.Logger.Fatal(err)
		}
	}()

	wg.Wait()

	// Run Seeders
	// factories.Factory()
}
