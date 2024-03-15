package main

import (
	"testing"
)

func TestIsValidLuhn(t *testing.T) {
	// Test Case 1: Valid Card Number
	cardNumber := "4111111111111111"
	result := isValidLuhn(cardNumber)
	expected := true

	if result != expected {
		t.Errorf("Expected %t, but got %t", expected, result)
	}

	// Test Case 2: Invalid Card Number
	cardNumber = "4111111111111112"
	result = isValidLuhn(cardNumber)
	expected = false

	if result != expected {
		t.Errorf("Expected %t, but got %t", expected, result)
	}

	// Test Case 3: Remove symbols characters
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

	// Test Case 5: Remove spaces characters
	cardNumber = "4111  1111  1111  1111"
	result = isValidLuhn(cardNumber)
	expected = true

	if result != expected {
		t.Errorf("Expected %t, but got %t", expected, result)
	}

	// Test Case 6: Remove non-numeric
	cardNumber = "4111   1111abc1111/!#1111"
	result = isValidLuhn(cardNumber)
	expected = true

	if result != expected {
		t.Errorf("Expected %t, but got %t", expected, result)
	}
}
