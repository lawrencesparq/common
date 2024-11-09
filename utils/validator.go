package utils

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ValidateString(s, name string) error {
	if s == "" {
		return fmt.Errorf("%s cannot be empty", name)
	}
	return nil
}

func ValidateTpin(s, name string) error {
	digitsReg, _ := regexp.Compile(`^[0-9]*$`)
	if err := ValidateString(s, name); err != nil {
		return err
	}

	if !digitsReg.MatchString(s) {
		return fmt.Errorf("%s should contain numbers only", name)
	}

	if len(s) < 10 || len(s) > 10 {
		return fmt.Errorf("%s should be 10 characters long", name)
	}

	return nil
}

func ValidateNumber(s, name string) error {
	digitsReg, _ := regexp.Compile(`^[0-9]*$`)
	if err := ValidateString(s, name); err != nil {
		return err
	}

	if !digitsReg.MatchString(s) {
		return fmt.Errorf("%s should contain numbers only", name)
	}

	return nil
}

func ValidateStringReg(s, name string) error {
	digitsReg, _ := regexp.Compile(`^[A-Za-z]*$`)
	if err := ValidateString(s, name); err != nil {
		return err
	}

	if !digitsReg.MatchString(s) {
		return fmt.Errorf("%s should contain text only", name)
	}

	return nil
}

func ValidateEmail(s, name string) error {
	emailReg, _ := regexp.Compile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if err := ValidateString(s, name); err != nil {
		return err
	}

	if !emailReg.MatchString(s) {
		return fmt.Errorf("%s doesn't match %s format", name, name)
	}

	return nil
}

func ValidateContact(s, name string) error {
	phoneNumberReg, _ := regexp.Compile(`^0(95|96|97|76|77|75)[0-9]{7}$`)
	if err := ValidateString(s, name); err != nil {
		return err
	}

	if !phoneNumberReg.MatchString(s) {
		return fmt.Errorf("%s doesn't match %s format", name, name)
	}

	return nil
}

func ValidateNRC(s, name string) error {
	nrcReg, _ := regexp.Compile(`^(ZN[0-9]{6}|[0-9]{6}/[0-9]{2}/(1|2|3){1})$`)
	if err := ValidateString(s, name); err != nil {
		return err
	}

	if !nrcReg.MatchString(s) {
		return fmt.Errorf("%s doesn't match %s format", name, name)
	}

	return nil
}
