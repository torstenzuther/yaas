package model

import (
	"time"

	"gorm.io/gorm"
)

type Grant struct {
	gorm.Model
	GrantValue          string    `gorm:"not null"`
	GrantType           string    `gorm:"not null"`
	ExpiresAt           time.Time `gorm:"not null"`
	CodeChallenge       string
	CodeChallengeMethod string `gorm:"not null"`
	RedirectUri         string `gorm:"not null"`
	Scope               string
	UserID              uint `gorm:"not null"`
	ClientID            uint `gorm:"not null"`
	Invalidated         bool `gorm:"not null"`
}
