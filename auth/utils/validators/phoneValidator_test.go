package validators

import (
	"testing"
)

func TestIsValidPhone(t *testing.T) {
	testCases := []struct {
		phone    string
		expected bool
	}{
		{"+1234567890123", true},   // Length greater than 13
		{"+123456789012", true},    // Valid phone number
		{"+12345678901", true},     // Length less than 13
		{"+12345678901a", false},   // Contains non-numeric character
		{"+12345678901", true},     // Length less than 13
		{"+12", false},             // Length less than 13
		{"123456789012", false},    // Missing '+'
		{"++123456789012", false},  // Extra '+'
		{"+1234567890a2", false},   // Non-numeric in main part
		{"+12345678901234", false}, // Length greater than 13
		{"+1", false},              // Length less than 13
	}

	for _, testCase := range testCases {
		result := IsValidPhone(testCase.phone)
		if result != testCase.expected {
			t.Errorf("IsValidPhone(%q) = %v; expected %v", testCase.phone, result, testCase.expected)
		}
	}
}
