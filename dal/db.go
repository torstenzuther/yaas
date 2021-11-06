package dal

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"yaas/dal/model"
)

func initDatabase(connection string) (*gorm.DB, error) {
	var db, err = gorm.Open(sqlite.Open(connection), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&model.User{}, &model.Client{}, &model.RedirectUri{}, &model.Grant{}); err != nil {
		return nil, err
	}
	return db, nil
}

func InitDAL(sessionStoreCleanupPeriod time.Duration, maxAge int, keys ...[]byte) (*Stores, chan struct{}, error) {
	db, err := initDatabase("gorm.db")
	if err != nil {
		return nil, nil, err
	}
	store, quit := initSessionStore(db, sessionStoreCleanupPeriod, maxAge, keys...)
	return &Stores{
		SessionStore: store,
		GrantStore:   newGrantStore(db),
		ClientStore:  newClientStore(db),
		UserStore:    newUserStore(db),
	}, quit, nil
}
