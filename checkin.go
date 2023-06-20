package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/smartclash/dailychecker/config"
	"os"
	"time"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Could not load .env file")
	}
}

func StartDailyCheckin() {
	members := GetMembersFromCheckinRole()

	for _, member := range members {
		CreateCheckinThreadForMember(member)
		time.Sleep(500 * time.Millisecond)
	}
}

func GetMembersFromCheckinRole() (checkinMembers []*discordgo.Member) {
	members, err := config.Discord.GuildMembers(os.Getenv("GUILD_ID"), "", 100)
	if err != nil {
		fmt.Println("Could not get list of guild members", err)
	}

	checkinRoleID := os.Getenv("ROLE_ID")
	for _, member := range members {
		for _, role := range member.Roles {
			if role == checkinRoleID {
				checkinMembers = append(checkinMembers, member)
			}
		}
	}

	return
}

func CreateCheckinThreadForMember(member *discordgo.Member) {
	username := member.Nick
	if username == "" {
		username = member.User.Username
	}

	thread, err := config.Discord.ThreadStartComplex(
		os.Getenv("CHECKIN_CHANNEL"),
		&discordgo.ThreadStart{
			Name:                "Daily Checkin for " + username,
			Type:                discordgo.ChannelTypeGuildPublicThread,
			AutoArchiveDuration: 15,
		},
	)
	if err != nil {
		fmt.Println("Could not create a thread", err)
	}

	if err = config.Discord.ThreadMemberAdd(thread.ID, member.User.ID); err != nil {
		fmt.Println("Could not add", member.User.Username, "to thread")
	}

	message := "Hiya, I am your daily checker. What did you do yesterday at work?"
	_, err = config.Discord.ChannelMessageSend(thread.ID, message)
	if err != nil {
		fmt.Println("Couldn't send a message inside thread")
	}
}
