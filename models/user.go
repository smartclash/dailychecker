package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string
	DiscordID string

	Responses []Response
}
