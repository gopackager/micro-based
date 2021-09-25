package helper

import (
	"os"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Env(v string, fallback ...string) string {
	val := os.Getenv(v)
	if len(fallback) == 0 {
		return val
	}
	if val == "" || len(fallback) > 0 {
		return fallback[0]
	}
	return val
}

func GenerateFromPassword(val []byte) []byte {
	hash, err := bcrypt.GenerateFromPassword(val, bcrypt.MinCost)
	if err != nil {
		return nil
	}
	return hash
}

func CompareHashAndPassword(hashed []byte, pwd []byte) bool {
	if err := bcrypt.CompareHashAndPassword(hashed, pwd); err != nil {
		return false
	}
	return true
}

func GenerateUUID() string {
	return uuid.New().String()
}

func InArray(value string, needed []string) bool {
	for _, v := range needed {
		if v == value {
			return true
		}
	}

	return false
}
