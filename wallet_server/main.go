package main

import (
	"os"
	"strconv"

	//"github.com/urfave/cli"
	"goblockchain/wallet_server/wallet_server"
	"log"
)

func init() {
	log.SetPrefix("Wallet Server:")
}

func main() {
	//port := flag.Uint("port", 8080, "TCP Port NUmber for Wallet Server")
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	portInt, _ := strconv.Atoi(port)

	gateway, ok := os.LookupEnv("GATEWAY")
	if !ok {
		gateway = "https://blockchain-server1:5000"
	}

	app := wallet_server.NewWalletServer(uint16(portInt), gateway)
	app.Run()

	//cliApp := cli.NewApp()
	//cliApp.Name = "wallet"
	//cliApp.Usage = "walletサーバー起動"
	//cliApp.Version = "0.0.1"
	//cliApp.Action = func(c *cli.Context) error {
	//	port := 8080
	//	if c.Bool("port") {
	//		port, _ = strconv.Atoi(c.Args().Get(0))
	//	}
	//	gateway := "http://127.0.0.1:5000"
	//	if c.Bool("gateway") {
	//		gateway = c.Args().Get(0)
	//	}
	//	app := wallet_server.NewWalletServer(uint16(port), gateway)
	//	app.Run()
	//	return nil
	//}
	//cliApp.Flags = []cli.Flag{
	//	&cli.Int64Flag{
	//		Name:  "port",
	//		Usage: "ポート番号",
	//	},
	//}

	//if err := cliApp.Run(os.Args); err != nil {
	//	return
	//}
}
