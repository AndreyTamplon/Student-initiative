package SQLStore

import (
	"APIServerSI/model/petition"
	"APIServerSI/store"
	"database/sql"
	"github.com/lib/pq"
)

type PetitionRepository struct {
	store *Store
}

func (repository *PetitionRepository) Create(petition *petition.Petition) error {
	if err := petition.Validate(); err != nil {
		return err
	}
	if err := repository.store.db.QueryRow(
		"INSERT INTO petitions (title, author_name, author_email, date_of_creation, date_of_expiration, tags, petition_content, number_of_signatures, signatures_target) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id",
		petition.Title, petition.AuthorName, petition.AuthorEmail, petition.DateOfCreation, petition.DateOfExpiration, pq.Array(petition.Tags), petition.PetitionContent, petition.NumberOfSignatures, petition.SignaturesTarget,
	).Scan(&petition.ID); err != nil {
		return err
	}
	return nil
}

func (repository *PetitionRepository) FindByID(id int) (*petition.Petition, error) {
	newPetition := &petition.Petition{}
	if err := repository.store.db.QueryRow(
		"SELECT id, title, author_name, author_email, date_of_creation, date_of_expiration, tags, petition_content, number_of_signatures, signatures_target FROM petitions WHERE id = $1",
		id,
	).Scan(
		&newPetition.ID, &newPetition.Title, &newPetition.AuthorName, &newPetition.AuthorEmail, &newPetition.DateOfCreation, &newPetition.DateOfExpiration, &newPetition.Tags, &newPetition.PetitionContent, &newPetition.NumberOfSignatures, &newPetition.SignaturesTarget,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return newPetition, nil
}

func (repository *PetitionRepository) FindByEmail(email string) (*petition.Petition, error) {
	newPetition := &petition.Petition{}
	if err := repository.store.db.QueryRow(
		"SELECT id, title, author_name, author_email, date_of_creation, date_of_expiration, tags, petition_content, number_of_signatures, signatures_target FROM petitions WHERE author_email = $1",
		email,
	).Scan(
		&newPetition.ID, &newPetition.Title, &newPetition.AuthorName, &newPetition.AuthorEmail, &newPetition.DateOfCreation, &newPetition.DateOfExpiration, &newPetition.Tags, &newPetition.PetitionContent, &newPetition.NumberOfSignatures, &newPetition.SignaturesTarget,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return newPetition, nil
}

func (repository *PetitionRepository) UpdatePetition(petition *petition.Petition) error {
	if err := petition.Validate(); err != nil {
		return err
	}
	if err := repository.store.db.QueryRow(
		"UPDATE petitions SET title = $1, author_name = $2, author_email = $3, date_of_creation = $4, date_of_expiration = $5, tags = $6, petition_content = $7, number_of_signatures = $8, signatures_target = $9 WHERE id = $10 RETURNING id",
		petition.Title, petition.AuthorName, petition.AuthorEmail, petition.DateOfCreation, petition.DateOfExpiration, petition.Tags, petition.PetitionContent, petition.NumberOfSignatures, petition.SignaturesTarget, petition.ID,
	).Scan(&petition.ID); err != nil {
		return err
	}
	return nil
}

func (repository *PetitionRepository) DeletePetition(id int) error {
	if err := repository.store.db.QueryRow(
		"DELETE FROM petitions WHERE id = $1",
		id,
	).Scan(); err != nil {
		return err
	}
	return nil
}

func (repository *PetitionRepository) FindByTitle(title string) (*petition.Petition, error) {
	petition := &petition.Petition{}
	if err := repository.store.db.QueryRow(
		"SELECT id, title, author_name, author_email, date_of_creation, date_of_expiration, tags, petition_content, number_of_signatures, signatures_target FROM petitions WHERE title = $1",
		title,
	).Scan(
		&petition.ID, &petition.Title, &petition.AuthorName, &petition.AuthorEmail, &petition.DateOfCreation, &petition.DateOfExpiration, pq.Array(&petition.Tags), &petition.PetitionContent, &petition.NumberOfSignatures, &petition.SignaturesTarget,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return petition, nil
}

func (repository *PetitionRepository) GetPetitionsByIdRange(idFrom int, idTo int) ([]petition.Petition, error) {
	var petitions []petition.Petition
	rows, err := repository.store.db.Query(
		"SELECT id, title, author_name, author_email, date_of_creation, date_of_expiration, tags, petition_content, number_of_signatures, signatures_target FROM petitions WHERE id BETWEEN $1 AND $2",
		idFrom, idTo,
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		newPetition := petition.Petition{}
		if err := rows.Scan(
			&newPetition.ID, &newPetition.Title, &newPetition.AuthorName, &newPetition.AuthorEmail, &newPetition.DateOfCreation, &newPetition.DateOfExpiration, pq.Array(&newPetition.Tags), &newPetition.PetitionContent, &newPetition.NumberOfSignatures, &newPetition.SignaturesTarget,
		); err != nil {
			return nil, err
		}
		petitions = append(petitions, newPetition)
	}
	return petitions, nil
}

func (repository *PetitionRepository) SignPetition(userId int, petitionId int) error {
	var ret []int64
	if err := repository.store.db.QueryRow(
		"SELECT signatories FROM petitions WHERE id = $1", petitionId).Scan(
		pq.Array(&ret),
	); err != nil {
		return err
	}
	for _, signatory := range ret {
		if int(signatory) == userId {
			return store.ErrAlreadySigned
		}
	}

	if _, err := repository.store.db.Exec(
		"UPDATE petitions SET number_of_signatures = number_of_signatures + 1, signatories = array_append(signatories, $2) WHERE id = $1", petitionId, userId); err != nil {
		return err
	}
	return nil
}
