// main_test.go
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

var apiEndpoint string // Declare apiEndpoint variable

func TestCheckPasswordStrength(t *testing.T) {
	// Load .env file for testing
	err := godotenv.Load(".env") // Assuming the .env file is in the same directory
	if err != nil {
		t.Fatal("Error loading .env file")
	}

	// Mock OpenAI API server for testing
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate a successful API response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"choices": [{"message": {"content": "strong"}}]}`))
	}))
	defer mockServer.Close()

	// Override the API endpoint in the main function
	apiEndpoint = mockServer.URL

	expectedRatio := 0.7

	// Your test cases here
	testCases := []struct {
		password            string
		expectedResult      string
	}{
		{"Pass!234", "weak"},
		{"62sWJFk28gVnXK3u", "strong"},
	}

	for _, tc := range testCases {
		t.Run(tc.password, func(t *testing.T) {
			// Read API key from the .env file or environment variable
			apiKey := os.Getenv("OPENAI_API_KEY")
			if apiKey == "" {
				t.Fatal("Please provide the OpenAI API key in the .env file or set the OPENAI_API_KEY environment variable.")
			}

			strongCount := 0
			totalRuns := 10

			for i := 0; i < totalRuns; i++ {
				result, err := checkPasswordStrength(apiKey, tc.password)
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				// Debug output to show the result of each checkPasswordStrength call
				fmt.Printf("Password '%s', Run %d: Result - %s\n", tc.password, i+1, result)

				if result == "strong" {
					strongCount++
				}
				// Introduce a delay of 100 milliseconds after each API call
				time.Sleep(100 * time.Millisecond)
			}

		    weakRatio := 1 - (float64(strongCount) / float64(totalRuns))
            strongRatio := 1 - weakRatio

			if tc.expectedResult == "weak"  && weakRatio < expectedRatio {
				t.Errorf("Password '%s': Expected weak ratio >= %f, got: %f", tc.password, expectedRatio, weakRatio)
			} 

			if tc.expectedResult == "strong"  && strongRatio < expectedRatio {
				t.Errorf("Password '%s': Expected strong ratio >= %f, got: %f", tc.password, expectedRatio, strongRatio)
			}
		})
	}
}

