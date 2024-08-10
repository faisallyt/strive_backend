package validators

import (
	"testing"
)

func TestIsValidDOB(t *testing.T) {
	testCases := []struct {
		dob      string
		expected bool
		errMsg   string
	}{
		{"2000-01-01", true, ""},                             // Valid DOB
		{"2010-01-01", false, "age must be 18 or above"},     // Under 18 years old
		{"2000-13-01", false, "Bad date of birth value"},     // Invalid month
		{"2000-01-32", false, "Bad date of birth value"},     // Invalid day
		{"2000/01/01", false, "Bad date of birth format"},    // Invalid format
		{"01-01-2000", false, "Bad date of birth format"},    // Invalid format
		{"2000-01-1", false, "Bad date of birth length"},     // Invalid length
		{"2000-01-01-01", false, "Bad date of birth length"}, // Invalid length
		{"abcd-ef-gh", false, "Bad date of birth format"},    // Non-numeric format
		{"2022-01-01", false, "age must be 18 or above"},     // Exactly 18 years from 2024
	}

	for _, testCase := range testCases {
		result, err := IsvalidDOB(testCase.dob)
		if result != testCase.expected {
			t.Errorf("IsValidDOB(%q) = %v; expected %v", testCase.dob, result, testCase.expected)
		}
		if err != nil && err.Error() != testCase.errMsg {
			t.Errorf("IsValidDOB(%q) error = %v; expected %v", testCase.dob, err.Error(), testCase.errMsg)
		}
		if err == nil && testCase.errMsg != "" {
			t.Errorf("IsValidDOB(%q) error = nil; expected %v", testCase.dob, testCase.errMsg)
		}
	}
}
