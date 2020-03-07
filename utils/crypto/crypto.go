package crypto

import "golang.org/x/crypto/bcrypt"

// ConvertToByte converts a string to a byte
func ConvertToByte(password string) []byte {
	return []byte(password)
}

// GenerateHashedPassowrd generates a bcrypt password that is hashed,
// if there is no error, it converts the byte to a string, and returns a nil error
// otherwiser it return and empty sting and the error
func GenerateHashedPassowrd(password string) (hashed string, err error) {
	b := ConvertToByte(password)
	hash, err := bcrypt.GenerateFromPassword(b, 12)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// VerifyPassword takes the hashed password and plain password submitte and verifys them.
// Returns true if they match, or false if not
func VerifyPassword(hahsed string, password string) bool {
	hash := ConvertToByte(hahsed)
	b := ConvertToByte(password)

	// compare passwords, if they don't match, err will not be nil
	err := bcrypt.CompareHashAndPassword(hash, b)
	if err != nil {
		return false
	}

	return true
}
