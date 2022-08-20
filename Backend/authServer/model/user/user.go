package user

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type User struct {
	ID                        int    `json:"id"`
	Name                      string `json:"name"`
	Email                     string `json:"email"`
	Confirmed                 bool   `json:"confirmed"`
	EncryptedConfirmationCode string `json:"encrypted_confirmation_code"`
	Password                  string `json:"password,omitempty"`
	EncryptedPassword         string `json:"-"`
}

var validDomains = []string{"g.nsu.ru", "alumni.nsu", "nsu.ru"}

func (user *User) BeforeCreate() error {
	if len(user.Password) > 0 {
		enc, err := encryptString(user.Password)
		if err != nil {
			return err
		}
		user.EncryptedPassword = enc
	}
	return nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (user *User) Validate() error {

	err := validation.ValidateStruct(
		user,
		validation.Field(&user.Name, validation.Required, validation.Length(1, 50)),
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.By(requiredIf(user.EncryptedPassword == "")), validation.Length(6, 100)),
	)
	if err == nil {
		if !user.checkEmailDomain() {
			err = errors.New("invalid_domain")
		}
	}
	return err

}

func (user *User) Sanitize() {
	user.Password = ""
}

func (user *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password)) == nil
}

func (user *User) checkEmailDomain() bool {
	emailDomain := user.Email[strings.LastIndex(user.Email, "@")+1:]
	for _, validDomain := range validDomains {
		if emailDomain == validDomain {
			return true
		}
	}
	return false
}
