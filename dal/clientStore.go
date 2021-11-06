package dal

import (
	"gorm.io/gorm"
	"yaas/dal/model"
)

type (
	ClientStore interface {
		GetClient(clientId string) (*Client, error)
	}

	Client struct {
		ClientID     string
		IsPublic     bool
		AllowHTTP    bool
		RedirectURIs []string
	}

	gormClientStore struct {
		db *gorm.DB
	}
)

func (clientStore *gormClientStore) GetClient(clientId string) (*Client, error) {
	var client model.Client
	if tx := clientStore.
		db.
		Where(&model.Client{ClientID: clientId}).
		Preload("RedirectUris").
		First(&client); tx.Error != nil {
		return nil, tx.Error
	}
	return &Client{
		ClientID:     client.ClientID,
		IsPublic:     client.IsPublic,
		AllowHTTP:    false, // TODO only allow localhost for HTTP?
		RedirectURIs: client.RedirectUris.Uris(),
	}, nil
}

func newClientStore(db *gorm.DB) *gormClientStore {
	return &gormClientStore{
		db: db,
	}
}
