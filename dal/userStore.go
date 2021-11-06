package dal

import (
	"gorm.io/gorm"
	"yaas/dal/model"
	"yaas/logic"
)

type (
	UserStore interface {
		Authenticate(username string, password string) bool
	}
)

type (
	gormUserStore struct {
		db *gorm.DB
	}
)

func (store *gormUserStore) Authenticate(username string, password string) bool {
	var user model.User
	err := store.db.Where(&model.User{UserName: username}).First(&user).Error
	if err != nil {
		return false
	}
	passwordHash := logic.HashSecret([]byte(password))
	if passwordHash != user.PasswordHash {
		return false
	}
	return true
}

func newUserStore(db *gorm.DB) *gormUserStore {
	return &gormUserStore{db: db}
}
