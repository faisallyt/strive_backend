package validators

import (
	"errors"
	"regexp"
)

func IsValidEmail(email string) (bool, error) {
	if len(email) >= 50 {
		return false, errors.New("Too long email address")
	}

	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return false, errors.New("Bad email format")
	}

	return true, nil

}
