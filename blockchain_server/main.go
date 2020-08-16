package main

import (
	"flag"
	"goblockchain/blockchain_server/blockchain_server"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	port := flag.Uint("port", 5000, "TCP Port Number for Blockchain Server")
	flag.Parse()
	app := blockchain_server.NewBlockchainServer(uint16(*port))
	app.Run()
}
