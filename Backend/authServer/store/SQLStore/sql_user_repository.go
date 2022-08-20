package SQLStore

import (
	"authorizationServer/model/user"
	"authorizationServer/store"
	"database/sql"
)

type UserRepository struct {
	store *Store
}

func (repository *UserRepository) Create(user *user.User) error {

	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.BeforeCreate(); err != nil {
		return err
	}

	if err := repository.store.db.QueryRow(
		"INSERT INTO users (name, email, confirmed, encrypted_confirmation_code, encrypted_password) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		user.Name, user.Email, user.Confirmed, user.EncryptedConfirmationCode, user.EncryptedPassword,
	).Scan(
		&user.ID,
	); err != nil {
		return err
	}
	return nil
}

func (repository *UserRepository) FindByEmail(email string) (*user.User, error) {
	user := &user.User{}
	if err := repository.store.db.QueryRow(
		"SELECT * FROM users WHERE email = $1",
		email,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Confirmed,
		&user.EncryptedConfirmationCode,
		&user.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return user, nil
}

func (repository *UserRepository) FindByID(id int) (*user.User, error) {
	user := &user.User{}
	if err := repository.store.db.QueryRow(
		"SELECT * FROM users WHERE id = $1",
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Confirmed,
		&user.EncryptedConfirmationCode,
		&user.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return user, nil
}

func (repository *UserRepository) DeleteUserByEmail(email string) error {
	if _, err := repository.store.db.Exec("DELETE FROM users WHERE email = $1", email); err != nil {
		return err
	}
	return nil
}

func (repository *UserRepository) UpdateUser(user *user.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.BeforeCreate(); err != nil {
		return err
	}
	if _, err := repository.store.db.Exec(
		"UPDATE users SET name = $1, email = $2, confirmed = $3, encrypted_confirmation_code = $4, encrypted_password = $5 WHERE id = $6",
		user.Name, user.Email, user.Confirmed, user.EncryptedConfirmationCode, user.EncryptedPassword, user.ID,
	); err != nil {
		return err
	}
	return nil
}
