package model

import "gorm.io/gorm"

type (
	RedirectUri struct {
		gorm.Model
		ClientID uint   `gorm:"not null; uniqueIndex:redirect_uri_unique"`
		Uri      string `gorm:"not null; uniqueIndex:redirect_uri_unique"`
	}

	RedirectUris []RedirectUri
)

func (r RedirectUris) Uris() []string {
	var uris []string
	for _, uri := range r {
		uris = append(uris, uri.Uri)
	}
	return uris
}
