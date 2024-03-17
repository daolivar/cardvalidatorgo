package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

type CardIssuer string

// Constants representing card issuers
const (
	AmericanExpress CardIssuer = "American Express"
	Discover        CardIssuer = "Discover"
	Mastercard      CardIssuer = "Mastercard"
	Visa            CardIssuer = "Visa"
	Unknown         CardIssuer = "Unknown"
	// Add more card networks as needed
)

// isValidLuhn checks if a given credit card number passes the Luhn algorithm
func isValidLuhn(cardNumber string) bool {
	// Remove all non-numeric characters
	cardNumbers := removeNonNumeric(cardNumber)

	// Convert the card number to a slice of digits
	digits := make([]int, len(cardNumbers))
	for i, char := range cardNumbers {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return false // Invalid character in card number
		}
		digits[i] = digit
	}

	// Iterate over the digits in reverse order, starting from the second-to-last digit
	for i := len(digits) - 2; i >= 0; i -= 2 {
		digits[i] *= 2     // Double the digit
		if digits[i] > 9 { // If the result is greater than 9, subtract 9
			digits[i] -= 9
		}
	}

	// Sum all the digits
	sum := 0
	for _, digit := range digits {
		sum += digit
	}

	// If the sum is a multiple of 10, the card number is valid
	return sum%10 == 0
}

func removeNonNumeric(s string) string {
	reg := regexp.MustCompile("[^0-9]")
	return reg.ReplaceAllString(s, "")
}

// getCardIssuer identifies the card issuer based on the card number prefix
func getCardIssuer(cardNumber string) CardIssuer {
	// Extract the first character of the card number to identify the prefix
	prefix := cardNumber[:1]

	// Check the prefix to determine the card issuer
	switch prefix {
	// American Express starts with '3' and has a second digit of '4' or '7'
	case "3":
		if cardNumber[1] == '4' || cardNumber[1] == '7' {
			return AmericanExpress
		}
	// Visa starts with '4' and has a length of either 13 or 16 digits
	case "4":
		if len(cardNumber) == 13 || len(cardNumber) == 16 {
			return Visa
		}
	// Mastercard starts with '5' and the second digit is between '1' and '5' (inclusive)
	case "5":
		if cardNumber[1] >= '1' && cardNumber[1] <= '5' {
			return Mastercard
		}
	// Discover starts with '6', and the second digit is '5', or the first four digits are '6011'
	case "6":
		if cardNumber[1] == '5' || cardNumber[:4] == "6011" {
			return Discover
		}
	}

	// If the prefix does not match known patterns, return "Unknown"
	return Unknown
}

// Handler handles incoming HTTP requests.
func handler(w http.ResponseWriter, r *http.Request) {
	// Implement handler logic
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	fmt.Printf("%t %v\n", isValidLuhn(""), getCardIssuer(""))
}
