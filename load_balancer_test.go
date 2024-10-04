package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
  "fmt"
  "crypto/sha256"
)

// hashString creates a SHA-256 hash of the input string
func hashString(input string) string {
	hash := sha256.Sum256([]byte(input))
	return fmt.Sprintf("%x", hash)
}

// TestHandleRequest tests the main request handler function
func TestHandleRequest(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Override the databases slice with test files
	databases = []string{
		tempDir + "/1db.txt",
		tempDir + "/2db.txt",
		tempDir + "/3db.txt",
		tempDir + "/4db.txt",
	}

	// Test cases
	testCases := []struct {
		name           string
		method         string
		body           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid POST request",
			method:         "POST",
			body:           "test data",
			expectedStatus: http.StatusOK,
			expectedBody:   "Request processed and written to",
		},
		{
			name:           "Invalid GET request",
			method:         "GET",
			body:           "",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Only POST requests are allowed",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(tc.method, "/", bytes.NewBufferString(tc.body))
			if err != nil {
				t.Fatalf("Could not create request: %v", err)
			}

			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()

			// Call the handler function
			handleRequest(rr, req)

			// Check the status code
			if status := rr.Code; status != tc.expectedStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, tc.expectedStatus)
			}

			// Check the response body
			if !strings.Contains(rr.Body.String(), tc.expectedBody) {
				t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), tc.expectedBody)
			}
		})
	}
}

// TestHashString tests the string hashing function
func TestHashString(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello", "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"},
		{"", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{"test data", "916f0027a575074ce72a331777c3478d6513f786a591bd892da1a577bf2335f9"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got := hashString(tc.input)
			if got != tc.expected {
				t.Errorf("hashString(%q) = %v; want %v", tc.input, got, tc.expected)
			}
		})
	}
}

// TestWriteToFile tests the file writing function
func TestWriteToFile(t *testing.T) {
	// Create a temporary file
	tmpfile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	// Test writing to the file
	content := "test content"
	err = writeToFile(tmpfile.Name(), content)
	if err != nil {
		t.Fatalf("Failed to write to file: %v", err)
	}

	// Read the file content
	fileContent, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	// Check if the content was written correctly
	if string(fileContent) != content+"\n" {
		t.Errorf("File content does not match: got %q, want %q", string(fileContent), content+"\n")
	}
}
