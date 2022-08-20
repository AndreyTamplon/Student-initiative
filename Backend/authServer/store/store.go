package store

type Store interface {
	Users() UserRepository
}
