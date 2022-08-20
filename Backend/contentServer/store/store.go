package store

type Store interface {
	Petitions() PetitionRepository
}
