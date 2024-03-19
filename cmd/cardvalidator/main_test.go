package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIsValidLuhn(t *testing.T) {
	// Test Case 1: Valid card number
	cardNumber := "4111111111111111"
	result := isValidLuhn(cardNumber)
	expected := true

	if result != expected {
		t.Errorf("Expected %t, but got %t", expected, result)
	}

	// Test Case 2: Invalid card number
	cardNumber = "4111111111111112"
	result = isValidLuhn(cardNumber)
	expected = false

	if result != expected {
		t.Errorf("Expected %t, but got %t", expected, result)
	}

	// Test Case 3: Remove symbol characters
	cardNumber = "4111-1111-1111-1111"
	result = isValidLuhn(cardNumber)
	expected = true

	if result != expected {
		t.Errorf("Expected %t, but got %t", expected, result)
	}

	// Test Case 4: Remove letters
	cardNumber = "4111a1111b1111c1111"
	result = isValidLuhn(cardNumber)
	expected = true

	if result != expected {
		t.Errorf("Expected %t, but got %t", expected, result)
	}

	// Test Case 5: Remove space characters
	cardNumber = "4111  1111  1111  1111"
	result = isValidLuhn(cardNumber)
	expected = true

	if result != expected {
		t.Errorf("Expected %t, but got %t", expected, result)
	}

	// Test Case 6: Remove all non-numeric characters
	cardNumber = "4111   1111abc1111/!#1111"
	result = isValidLuhn(cardNumber)
	expected = true

	if result != expected {
		t.Errorf("Expected %t, but got %t", expected, result)
	}
}

func TestGetCardIssuer(t *testing.T) {
	// Test Case 1: Visa card number
	cardNumber := "4111111111111111"
	result := getCardIssuer(cardNumber)
	expected := Visa

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Test Case 2: Mastercard card number
	cardNumber = "5555555555554444"
	result = getCardIssuer(cardNumber)
	expected = Mastercard

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Test Case 3: American Express card number
	cardNumber = "378282246310005"
	result = getCardIssuer(cardNumber)
	expected = AmericanExpress

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Test Case 4: Discover card number
	cardNumber = "6534567890123452"
	result = getCardIssuer(cardNumber)
	expected = Discover

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Test Case 5: Discover card number 2
	cardNumber = "6011000000000004"
	result = getCardIssuer(cardNumber)
	expected = Discover

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestHandler(t *testing.T) {
	// Create a request body with a valid card number
	requestBody := []byte(`{"cardnumber": "4111111111111111"}`)

	// Create a new HTTP request with the request body
	req, err := http.NewRequest("POST", "/validate", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a response recorder to record the response from the handler
	rr := httptest.NewRecorder()

	// Call the handler function with the mock request and response recorder
	handler(rr, req)

	// Check the HTTP status code of the response
	if rr.Code != http.StatusOK {
		t.Errorf("Handler returned unexpected status code: got %v, want %v", rr.Code, http.StatusOK)
	}

	// Decode the response body to get the validation result
	var response CardValidationResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	// Check the validation result
	if !response.Valid {
		t.Errorf("Unexpected validation result: got %v, want true", response.Valid)
	}

	// Check the issuer result
	if response.Issuer != "Visa" {
		t.Errorf("Unexpected issuer result: got %s, want Visa", response.Issuer)
	}
}
