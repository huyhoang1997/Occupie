package main

import (
	"fmt"
	"occupie/config"
	"occupie/db"
	"occupie/router"
	"sync"
)

const welcome = "Welcome to Occupie"

var (
	wg = sync.WaitGroup{}
)

func main() {
	fmt.Println("Loading....")
	// Init config
	config.NewApplicationConfig()
	wg.Add(1)
	// Init application
	go func() {
		go db.NewDbConnection(config.DEV)
		wg.Done()
	}()
	wg.Wait()
	appRoute := router.InitRouter()
	// Done initialize application
	err := appRoute.Run("3000")
	if err != nil {
		panic(err)
	}
	fmt.Println(welcome)
}
