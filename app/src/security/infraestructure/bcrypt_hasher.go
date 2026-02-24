package infraestructure

import (
	"golang.org/x/crypto/bcrypt"
	securityDomain"sharedup/app/src/security/domain"
)

type BCryptHasher struct{}

func NewBCryptHasher() securityDomain.PasswordHasher {
	return &BCryptHasher{}

}

func (b *BCryptHasher) Hash(plain string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(bytes), err
}

func (b *BCryptHasher) Compare(hash, plain string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain)) == nil
}
