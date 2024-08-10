package validators

import (
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	testCases := []struct {
		email    string
		expected bool
		errMsg   string
	}{
		//sample test cases
		{"sadaosd@askou.com", true, ""},               // Valid email
		{"asdvasd@asdasd", false, "Bad email format"}, // Invalid email

		// write all possible test cases
		{"asdasd.com", false, "Bad email format"},          // Invalid email
		{"asdasd@asdasd", false, "Bad email format"},       // Invalid email
		{"asdasd@asdasd.", false, "Bad email format"},      // Invalid email
		{"asdasd@asdasd.c", false, "Bad email format"},     // Invalid email
		{"asdasd@asdasd.c.", false, "Bad email format"},    // Invalid email
		{"asdasd@asdasd.c.o", false, "Bad email format"},   // Invalid email
		{"asdasd@asdasd.c.o.", false, "Bad email format"},  // Invalid email
		{"asdasd@asdasd.c.o.m", false, "Bad email format"}, // Invalid email
	}

	for _, testCase := range testCases {
		result, err := IsValidEmail(testCase.email)
		if result != testCase.expected {
			t.Errorf("IsValidDOB(%q) = %v; expected %v", testCase.email, result, testCase.expected)
			if err != nil && err.Error() != testCase.errMsg {
				t.Errorf("IsValidDOB(%q) error = %v; expected %v", testCase.email, err.Error(), testCase.errMsg)
			}
			if err == nil && testCase.errMsg != "" {
				t.Errorf("IsValidDOB(%q) error = nil; expected %v", testCase.email, testCase.errMsg)
			}
		}
	}
}
