package main

import (
	"flag"
	"log"
	"sync"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	port := flag.Uint("port", 5000, "TCP Port Number for Blockchain Server")
	flag.Parse()
	app := NewBlockchainServer(uint16(*port))
	go func() {
		defer wg.Done()
		app.Run()
	}()

	wg.Wait()
}
