package hash

import "golang.org/x/crypto/bcrypt"

type Password interface {
	String(str string) (string, error)
	CheckString(str, strCheck string) error
}

type Hash struct{}

func NewPasswordHasher() *Hash {
	return &Hash{}
}
func (h *Hash) String(str string) (string, error) {
	pw := []byte(str)
	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func (h *Hash) CheckString(str, strCheck string) error {
	err := bcrypt.CompareHashAndPassword([]byte(str), []byte(strCheck))
	if err != nil {
		return err
	}
	return nil
}
