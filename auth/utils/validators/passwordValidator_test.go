package validators

import (
	"testing"
)

func TestIsValidPassword(t *testing.T) {
	testCases := []struct {
		password string
		expected bool
	}{
		{"Password1!", true},        // Valid password
		{"password1!", false},       // No uppercase letter
		{"PASSWORD1!", false},       // No lowercase letter
		{"Password", false},         // No special character
		{"Password123", false},      // No special character
		{"P@ssw0rd", true},          // Valid password
		{"Short1!", false},          // Too short
		{"noupper123@", false},      // No uppercase letter
		{"NOLOWER123@", false},      // No lowercase letter
		{"NoSpecialChar123", false}, // No special character
		{"Valid!Password123", true}, // Valid password
	}

	for _, testCase := range testCases {
		result := IsValidPassword(testCase.password)
		if result != testCase.expected {
			t.Errorf("IsValidPassword(%q) = %v; expected %v", testCase.password, result, testCase.expected)
		}
	}
}
