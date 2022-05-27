package bpArgon2

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/crypto/argon2"
)

// GeneratePassword is used to generate a new password hash for storing and
// comparing at a later date.
func GeneratePassword(password string) (string, error) {
	p := NewConfig()

	if len(password) <= 0 {
		return "", errors.New("empty password")
	}
	if len(password) < 8 {
		return "", errors.New("password too short")
	}
	// Generate a Salt
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, p.time, p.memory, p.threads, p.keyLen)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	format := "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
	full := fmt.Sprintf(format, argon2.Version, p.memory, p.time, p.threads, b64Salt, b64Hash)
	return full, nil
}

// ComparePassword is used to compare a user-inputted password to a hash to see
// if the password matches or not.
func ComparePassword(password, hash string) (bool, error) {

	parts := strings.Split(hash, "$")

	p := &PasswordConfig{}
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &p.memory, &p.time, &p.threads)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}
	p.keyLen = uint32(len(decodedHash))

	comparisonHash := argon2.IDKey([]byte(password), salt, p.time, p.memory, p.threads, p.keyLen)
	var compare = subtle.ConstantTimeCompare(decodedHash, comparisonHash) == 1
	if compare == false {
		return false, errors.New("Not valid password")
	}
	return compare, nil
}
