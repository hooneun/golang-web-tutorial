package helpers

import "golang.org/x/crypto/bcrypt"

// HashPassword !
func HashPassword(password *string) error {
	bytes := []byte(*password)
	hash, err := bcrypt.GenerateFromPassword([]byte(bytes), 14)
	if err != nil {
		return err
	}
	*password = string(hash[:])
	return nil
}

// CheckPasswordHash !
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
