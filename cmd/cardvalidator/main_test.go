package main

import (
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
