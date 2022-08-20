package authentification

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"time"
)

func GenerateRandomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func EncryptCode(code int) string {
	if encryptedCode, err := bcrypt.GenerateFromPassword([]byte(strconv.Itoa(code)), bcrypt.MinCost); err == nil {
		return string(encryptedCode)
	} else {
		return ""
	}
}

func CompareCode(code string, encryptedCode string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(encryptedCode), []byte(code)); err == nil {
		return true
	} else {
		return false
	}
}
