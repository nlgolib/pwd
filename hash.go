package pwd

import "golang.org/x/crypto/bcrypt"

// Hash generates a bcrypt hash from the given password.
// It uses the default cost for hashing.
// Returns the hashed password as a string.
func Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// Verify checks if the given password matches the provided hash.
// It returns true if the password is correct, false otherwise.
func Verify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
