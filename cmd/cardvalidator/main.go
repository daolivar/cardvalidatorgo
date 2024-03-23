package main

import (
	"encoding/json"
	"log"
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

// CardValidationRequest defines the structure of the JSON request
type CardValidationRequest struct {
	CardNumber string `json:"cardnumber"`
}

// CardValidationResponse defines the structure of the JSON response
type CardValidationResponse struct {
	Valid  bool       `json:"valid"`
	Issuer CardIssuer `json:"issuer"`
	Error  string     `json:"error,omitempty"`
}

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
func getCardIssuer(cardNumber string, isValid bool) CardIssuer {

	if isValid {
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
	}

	// If the prefix does not match known patterns, return "Unknown"
	return Unknown
}

// handler handles incoming HTTP requests.
func handler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		// If not POST, return "Method not allowed" error
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the request body into CardValidationRequest struct
	var request CardValidationRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Error decoding JSON data", http.StatusBadRequest)
		return
	}

	// Extract the card number from the decoded struct
	cardNumber := request.CardNumber

	// Validate the card number using the Luhn algorithm
	isValid := isValidLuhn(cardNumber)

	// Get the card issuer based on the card number prefix
	issuer := getCardIssuer(cardNumber, isValid)

	// Prepare the response object
	response := CardValidationResponse{
		Valid:  isValid,
		Issuer: issuer,
	}

	// Set the content type header to indicate JSON response
	w.Header().Set("Content-Type", "application/json")

	// Encode the response object to JSON and write it as the response
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Register a handler for the /validate endpoint
	http.HandleFunc("/validate", handler)

	// Start the server and listen for incoming connections
	log.Println("Server is listening on port 8080...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
