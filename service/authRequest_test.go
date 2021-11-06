package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAuthRequest(t *testing.T) {
	type testCase struct {
		params         map[string][]string
		expectedParams map[string]string
		expectedError  string
	}
	for _, test := range []testCase{
		{
			params: map[string][]string{
				"abc": {"123"},
			},
			expectedParams: map[string]string{},
		},
		{
			params: map[string][]string{
				"abc": {"123", "456"},
			},
			expectedParams: map[string]string{},
		},
		{
			params: map[string][]string{
				responseType: {"123"},
			},
			expectedParams: map[string]string{
				responseType: "123",
			},
		},
		{
			params: map[string][]string{
				responseType: {"code"},
				scope:        {"openid profile"},
			},
			expectedParams: map[string]string{
				responseType: "code",
				scope:        "openid profile",
			},
		},
		{
			params: map[string][]string{
				responseType: {"123", "456"},
			},
			expectedParams: map[string]string{},
			expectedError:  invalidRequest,
		},
	} {
		actual, err := parseAuthRequest(test.params)
		if test.expectedError == "" {
			assert.Nil(t, err)
			assert.EqualValues(t, actual.parameters, test.expectedParams)
		} else {
			assert.EqualError(t, err, test.expectedError)
			assert.Nil(t, actual)
		}
	}
}
