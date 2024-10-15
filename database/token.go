package database

import "templates/model"

type TokenRepository interface {
	UpsertToken(token *model.Token) (err error)
	GetTokenByNameCollection(name string) (token *model.Token, err error)
	DeleteToken(collection string) (err error)
}
