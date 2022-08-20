package store

import (
	"APIServerSI/model/petition"
)

type PetitionRepository interface {
	Create(petition *petition.Petition) error
	FindByID(int) (*petition.Petition, error)
	FindByTitle(string) (*petition.Petition, error)
	UpdatePetition(petition *petition.Petition) error
	DeletePetition(int) error
	GetPetitionsByIdRange(int, int) ([]petition.Petition, error)
	SignPetition(int, int) error
}
