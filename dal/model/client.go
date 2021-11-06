package model

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	ClientID                string `gorm:"unique;not null"`
	ClientSecret            string
	ClientType              string `gorm:"not null"`
	TokenEndpointAuthMethod string `gorm:"not null; default:client_secret_basic"`
	GrantTypes              string `gorm:"not null"`
	ResponseTypes           string `gorm:"not null"`
	ClientName              string
	ClientUri               string
	LogoUri                 string
	Scope                   string
	Contacts                string
	TosUri                  string
	PolicyUri               string
	JwksUri                 string
	Jwks                    string
	SoftwareId              string
	SoftwareVersion         string
	IsPublic                bool `gorm:"not null; default:0"`
	Disabled                bool `gorm:"not null; default:0"`
	AllowOfflineAccess      bool `gorm:"not null; default:0"`
	AccessTokenLifetime     int  `gorm:"not null; default:3600"`
	RedirectUris            RedirectUris
	Grants                  []Grant
}
