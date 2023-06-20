package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/smartclash/dailychecker/config"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Could not load .env file")
	}
}

func main() {
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	err := config.Discord.Close()
	if err != nil {
		fmt.Println("Could not close discord session")
	}
}
