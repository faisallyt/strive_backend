package validators

import (
	"errors"
	"regexp"
	"time"
)

func IsvalidDOB(dob string) (bool, error) {
	if len(dob) != 10 {
		return false, errors.New("Bad date of birth length")
	}

	re := regexp.MustCompile(`^[0-9]{4}-[0-9]{2}-[0-9]{2}$`)
	if !re.MatchString(dob) {
		return false, errors.New("Bad date of birth format")
	}

	// create a datetime object IN GO
	dateObj, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return false, errors.New("Bad date of birth value")
	}

	day := dateObj.Day()
	month := dateObj.Month()

	// get age
	age := time.Now().Year() - dateObj.Year()

	if time.Now().Month() < month || time.Now().Month() == month && time.Now().Day() < day {
		age--
	}

	if age < 18 {
		return false, errors.New("age must be 18 or above")
	}

	return true, nil

}

