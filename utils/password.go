package utils

import (
	"fmt"
	"strconv"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePasswords(hashedPassword string, plain []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), plain)
	return err == nil
}

func IsValidPassword(s string, min int) error {
	var (
		isMin   bool
		special bool
		number  bool
		upper   bool
		lower   bool
	)

	//test for the minimum characters required for the password string
	if len(s) < min {
		isMin = false
		return fmt.Errorf("length should be atleast " + strconv.Itoa(min) + " characters")
	}

	for _, c := range s {
		// Optimize perf if all become true before reaching the end
		if special && number && upper && lower && isMin {
			break
		}

		// else go on switching
		switch {
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsLower(c):
			lower = true
		case unicode.IsNumber(c):
			number = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		}
	}

	// Add custom error messages
	if !special {
		return fmt.Errorf("should contain at least a single special character")
	}
	if !number {
		return fmt.Errorf("should contain at least a single digit")
	}
	if !lower {
		return fmt.Errorf("should contain at least a single lowercase letter")
	}
	if !upper {
		return fmt.Errorf("should contain at least single uppercase letter")
	}

	// everyting is right
	return nil
}
