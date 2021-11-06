package logic

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCodeChallenge(t *testing.T) {
	type testCase struct {
		codeVerifier          string
		method                CodeMethod
		expectedCodeChallenge string
		expectedError         error
	}

	for _, test := range []testCase{
		{
			codeVerifier:  "123123",
			method:        Plain,
			expectedError: errors.New("invalid code verifier"),
		},
		{
			codeVerifier:          "RK-Ng69oNo1yvHP.T5IX5Lnug-OcYriSGu6.Ct740j0.W-fhBz5ww6dKTKHLuq3oN-GeoRUc_JpEvmjpjfkjUjHUcd6dQXr_PTgBTf3-.iyOM_J_26JaSH6myCjOib0u",
			method:                S256,
			expectedCodeChallenge: "xCzLV2UzgTZpeWl1EpKLeHzL62IopSaIfM2isMAn1AQ",
		},
		{
			codeVerifier:          "RK-Ng69oNo1yvHP.T5IX5Lnug-OcYriSGu6.Ct740j0.W-fhBz5ww6dKTKHLuq3oN-GeoRUc_JpEvmjpjfkjUjHUcd6dQXr_PTgBTf3-.iyOM_J_26JaSH6myCjOib0u",
			method:                Plain,
			expectedCodeChallenge: "RK-Ng69oNo1yvHP.T5IX5Lnug-OcYriSGu6.Ct740j0.W-fhBz5ww6dKTKHLuq3oN-GeoRUc_JpEvmjpjfkjUjHUcd6dQXr_PTgBTf3-.iyOM_J_26JaSH6myCjOib0u",
		},
	} {
		codeChallenge, err := createCodeChallenge([]byte(test.codeVerifier), test.method)
		if test.expectedError == nil {
			assert.Equal(t, codeChallenge, []byte(test.expectedCodeChallenge))
			assert.Nil(t, err)
		} else {
			assert.EqualError(t, err, test.expectedError.Error())
			assert.Nil(t, codeChallenge)
		}
	}

}

func TestParseScope(t *testing.T) {
	type testCase struct {
		scope           string
		expected        []string
		expectedSuccess bool
	}

	for _, test := range []testCase{
		{
			scope:           "",
			expected:        nil,
			expectedSuccess: false,
		},
		{
			scope:           " ",
			expected:        nil,
			expectedSuccess: false,
		},
		{
			scope:           "openid",
			expected:        []string{"openid"},
			expectedSuccess: true,
		},
		{
			scope:           "openid ",
			expected:        nil,
			expectedSuccess: false,
		},
		{
			scope:           " openid",
			expected:        nil,
			expectedSuccess: false,
		},
		{
			scope:           "openid profile",
			expected:        []string{"openid", "profile"},
			expectedSuccess: true,
		},
		{
			scope:           "openid profile ",
			expected:        nil,
			expectedSuccess: false,
		},
	} {
		actual, actualSuccess := ParseScope(test.scope)
		assert.EqualValues(t, actual, test.expected)
		assert.Equal(t, actualSuccess, test.expectedSuccess)
	}
}
