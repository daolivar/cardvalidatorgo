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
	Visa            CardIssuer = "Visa"
	Mastercard      CardIssuer = "Mastercard"
	AmericanExpress CardIssuer = "American Express"
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
	_ = cardNumber
	// Implement get card issuer
	return CardIssuer("unknown")
}

// Handler handles incoming HTTP requests.
func handler(w http.ResponseWriter, r *http.Request) {
	// Implement handler logic
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	fmt.Printf("%t %v\n", isValidLuhn(""), getCardIssuer(""))
}
