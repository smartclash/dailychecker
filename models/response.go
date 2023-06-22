package models

import "gorm.io/gorm"

type Response struct {
	gorm.Model
	Response string

	UserID int
	User   User

	QuestionID int
	Question   Question
}
