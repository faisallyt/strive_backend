package validators

import (
	"testing"
)

func TestIsValidUsername(t *testing.T) {
	testCases := []struct {
		username string
		expected bool
	}{
		{"user123", true},                   // Valid username
		{"user_123", true},                  // Valid username with underscore
		{"user", false},                     // Too short
		{"user12345678901234567890", false}, // Too long
		{"User123", true},                   // Valid username with uppercase
		{"user@123", false},                 // Invalid character '@'
		{"user-123", false},                 // Invalid character '-'
		{"", false},                         // Empty username
		{"user!123", false},                 // Invalid character '!'
		{"user space", false},               // Invalid character ' '
	}

	for _, testCase := range testCases {
		result := IsValidUsername(testCase.username)
		if result != testCase.expected {
			t.Errorf("IsValidUsername(%q) = %v; expected %v", testCase.username, result, testCase.expected)
		}
	}
}
