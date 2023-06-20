package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/smartclash/dailychecker/config"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Could not load .env file")
	}
}

func main() {
	// Set to Indian time
	location, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		fmt.Println("Could not load timezone information")
	}

	// Run the checkin process only by 11:30 every morning
	cron := gocron.NewScheduler(location)
	_, err = cron.Every(1).Day().At("11:30").Do(func() {
		fmt.Println("Starting check-in process")
		StartDailyCheckin()
	})
	if err != nil {
		fmt.Println("Could not run cron job")
	}

	cron.StartAsync()
	fmt.Println("Scheduler is now running")

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	err = config.Discord.Close()
	if err != nil {
		fmt.Println("Could not close discord session")
	}
}
