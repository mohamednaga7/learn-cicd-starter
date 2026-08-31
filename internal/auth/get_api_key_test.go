package auth

import (
	"net/http"
	"testing"
)

type TestCase struct {
	Title                string
	Headers              http.Header
	ExpectedErrorMessage string
	ExpectedResult       string
}

func TestGetAPIKey(test *testing.T) {
	testsValues := [5]TestCase{
		{"No Authorization Header", http.Header{}, "no authorization header included", ""},
		{"Empty Authorization Header", http.Header{"Authorization": []string{""}}, "no authorization header included", ""},
		{"No Token Provided", http.Header{"Authorization": []string{"ApiKey"}}, "malformed authorization header", ""},
		{"Correct Format But Empty Token", http.Header{"Authorization": []string{"ApiKey  "}}, "malformed authorization header", ""},
		{"Correct Authorization Header", http.Header{"Authorization": []string{"ApiKey sampleToken"}}, "", "sampleToken"},
	}

	for _, testCase := range testsValues {
		test.Run(testCase.Title, func(t *testing.T) {
			sampleHeaders := testCase.Headers

			result, err := GetAPIKey(sampleHeaders)

			if err == nil {
				if testCase.ExpectedErrorMessage != "" {
					t.Error("Should throw error")
				}
			} else {
				if err.Error() != testCase.ExpectedErrorMessage {
					t.Errorf("Should throw correct error message \nExpected: %s\nGot: %s", testCase.ExpectedErrorMessage, err)
				}
			}

			if result != testCase.ExpectedResult {
				t.Errorf("Should return correct result \nExpected: %s\nGot: %s", testCase.ExpectedResult, result)
			}
		})
	}
}
