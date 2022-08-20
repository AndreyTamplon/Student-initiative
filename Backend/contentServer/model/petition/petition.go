package petition

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Petition struct {
	ID                 int      `json:"id"`
	Title              string   `json:"title"`
	AuthorName         string   `json:"authorName"`
	AuthorEmail        string   `json:"authorEmail"`
	DateOfCreation     string   `json:"dateOfCreation"`
	DateOfExpiration   string   `json:"dateOfExpiration"`
	Tags               []string `json:"tags"`
	PetitionContent    string   `json:"petitionContent"`
	NumberOfSignatures int      `json:"numberOfSignatures"`
	SignaturesTarget   int      `json:"signaturesTarget"`
}

func (petition *Petition) Validate() error {
	return validation.ValidateStruct(
		petition,
		validation.Field(&petition.Title, validation.Required, validation.Length(1, 150)),
		validation.Field(&petition.AuthorName, validation.Required, validation.Length(1, 50)),
		validation.Field(&petition.AuthorEmail, validation.Required, is.Email),
		validation.Field(&petition.DateOfCreation, validation.Required, validation.Length(1, 30)),
		validation.Field(&petition.DateOfExpiration, validation.Required, validation.Length(1, 30)),
		validation.Field(&petition.Tags, validation.Each(validation.Length(1, 100))),
		validation.Field(&petition.PetitionContent, validation.Required, validation.Length(1, 5000)),
		validation.Field(&petition.NumberOfSignatures, validation.Min(0)),
		validation.Field(&petition.SignaturesTarget, validation.Required, validation.Min(0)),
	)
}
