package api

import (
	"errors"
	"testing"
)

func TestUserIsEligible(t *testing.T) {
	var tests = []struct {
		email       string
		password    string
		age         int
		expectedErr error
	}{
		{
			email:       "test@example.com",
			password:    "password",
			age:         21,
			expectedErr: nil,
		},
		{
			email:       "",
			password:    "password",
			age:         21,
			expectedErr: errors.New("email can't be empty"),
		},
		{
			email:       "test@example.com",
			password:    "",
			age:         21,
			expectedErr: errors.New("password can't be empty"),
		},
		{
			email:       "test@example.com",
			password:    "password",
			age:         1,
			expectedErr: errors.New("age must be at least 18 years"),
		},
	}

	for _, tt := range tests {
		err := userIsEligible(tt.email, tt.password, tt.age)
		errString := ""
		expectedErrString := ""
		if err != nil {
			errString = err.Error()
		}

		if tt.expectedErr != nil {
			expectedErrString = tt.expectedErr.Error()
		}

		if errString != expectedErrString {
			t.Errorf("got %s, want %s", errString, expectedErrString)
		}
	}

}
