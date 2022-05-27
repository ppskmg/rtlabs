package bpArgon2

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgon2_newConfig(t *testing.T) {
	c := NewConfig()
	assert.IsType(t, c, &PasswordConfig{})
}

func TestArgon2_GeneratePassword(t *testing.T) {
	pas := "password"
	c, _ := GeneratePassword(pas)
	assert.IsType(t, c, "string")
	assert.NotNil(t, c)
	assert.NotEmpty(t, c)
	fmt.Println(c)
	assert.Equal(t, len(strings.Split(c, ",")), 3)
	assert.Equal(t, len(strings.Split(c, "$")), 6)
	assert.Equal(t, len(strings.Split(strings.Split(c, "$")[3], ",")), 3)
}

func TestArgon2_GeneratePasswordErr(t *testing.T) {
	testCases := []struct {
		name        string
		payload     string
		expectedErr interface{}
	}{
		{
			name:        "valid",
			payload:     "longPassword",
			expectedErr: nil,
		},
		{
			name:        "empty",
			payload:     "",
			expectedErr: "empty password",
		},
		{
			name:        "empty",
			payload:     "short",
			expectedErr: "password too short",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := GeneratePassword(tc.payload)
			fmt.Println(err)
			if err == nil {
				assert.Equal(t, tc.expectedErr, nil)
			} else {
				assert.Equal(t, tc.expectedErr, err.Error())
			}
		})
	}
}
func TestArgon2_ComparePassword(t *testing.T) {
	pas := "validPassword"
	hash, _ := GeneratePassword(pas)
	testCases := []struct {
		name     string
		payload  string
		expected bool
	}{
		{
			name:     "valid",
			payload:  pas,
			expected: true,
		},
		{
			name:     "invalid: another password",
			payload:  "invalidPassword",
			expected: false,
		},
		{
			name:     "invalid: empty",
			payload:  "",
			expected: false,
		},
		{
			name:     "invalid: space before",
			payload:  " validPassword",
			expected: false,
		},
		{
			name:     "invalid: space after",
			payload:  "validPassword ",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cp, _ := ComparePassword(tc.payload, hash)
			assert.Equal(t, tc.expected, cp)
		})
	}
}
