package SQLStore

import (
	"authorizationServer/store"
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	db             *sql.DB
	UserRepository *UserRepository
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (store *Store) Users() store.UserRepository {
	if store.UserRepository != nil {
		return store.UserRepository
	}
	store.UserRepository = &UserRepository{
		store: store,
	}
	return store.UserRepository
}
