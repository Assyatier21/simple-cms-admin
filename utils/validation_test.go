package utils

import "testing"

func TestIsValidAlphabet(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"abcdefg", true},
		{"ABCDEFG", true},
		{"abc123", false},
		{"abc!@#", false},
		{"", true},
	}

	for _, tc := range testCases {
		result := IsValidAlphabet(tc.input)
		if result != tc.expected {
			t.Errorf("For input %s, expected %t but got %t", tc.input, tc.expected, result)
		}
	}
}

func TestIsValidNumeric(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected bool
	}{
		{
			desc:     "valid numeric",
			input:    "1234567890",
			expected: true,
		},
		{
			desc:     "invalid numeric",
			input:    "abc",
			expected: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := IsValidNumeric(tC.input)
			if result != tC.expected {
				t.Errorf("Expected %t, but got %t", tC.expected, result)
			}
		})
	}
}

func TestIsValidAlphaNumeric(t *testing.T) {
	result := IsValidAlphaNumeric("abc123")
	if !result {
		t.Errorf("Expected true, but got %t", result)
	}

	result = IsValidAlphaNumeric("abc123#")
	if result {
		t.Errorf("Expected false, but got %t", result)
	}
}

func TestIsValidSlug(t *testing.T) {
	result := IsValidSlug("abc-123")
	if !result {
		t.Errorf("Expected true, but got %t", result)
	}

	result = IsValidSlug("abc 123")
	if result {
		t.Errorf("Expected false, but got %t", result)
	}
}
