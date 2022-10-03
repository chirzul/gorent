package libs

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPass), nil
}

func CheckPassword(dbPass, inputPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(inputPass))
	return err == nil
}
