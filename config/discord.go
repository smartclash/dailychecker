package config

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var Discord *discordgo.Session

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Could not load .env file")
	}

	ConnectDiscord()
}

func ConnectDiscord() {
	var err error

	Discord, err = discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		fmt.Println("Could not connect to discord gateway")
		return
	}

	err = Discord.Open()
	if err != nil {
		fmt.Println("error opening discord bot session")
		return
	}
}
