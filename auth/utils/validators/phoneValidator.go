package validators

import "regexp"

func IsValidPhone(phone string) bool {
	if len(phone) < 12 || len(phone) > 14 {
		return false
	}

	re := regexp.MustCompile(`^\+[0-9]{1,3}[0-9]{10}$`)
	return re.MatchString(phone)
}
