package SQLStore

import (
	"APIServerSI/store"
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	db                 *sql.DB
	PetitionRepository *PetitionRepository
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (store *Store) Petitions() store.PetitionRepository {
	if store.PetitionRepository != nil {
		return store.PetitionRepository
	}
	store.PetitionRepository = &PetitionRepository{
		store: store,
	}
	return store.PetitionRepository
}
