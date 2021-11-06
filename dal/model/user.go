package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName     string `gorm:"uniqueIndex;not null"`
	PasswordHash string `gorm:"not null"`
	Sub          string `gorm:"uniqueIndex;not null"`
	Grants       []Grant
}
