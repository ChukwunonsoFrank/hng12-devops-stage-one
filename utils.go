package main

import (
	"math"
	"fmt"
	"io"
	"net/http"
	"unicode"
)

func isNumberPrime(number int) bool {
	if number <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number % i == 0 {
			return false
		}
	}

	return true
}

func isNumberOdd(number int) bool {
	return number % 2 != 0
}

func isNumberEven(number int) bool {
	return number % 2 == 0
}

func isNumberArmstrong(number int) bool {
	// Convert the number to a string to easily get the number of digits
	numStr := fmt.Sprintf("%d", number)
	numDigits := len(numStr)

	// Initialize a variable to store the sum of the cubes of the digits
	sum := 0

	// Iterate over each digit in the number
	for _, digit := range numStr {
		// Convert the digit back to an integer and calculate its cube
		digitCube := int(math.Pow(float64(digit-'0'), float64(numDigits)))
		// Add the cube to the sum
		sum += digitCube
	}

	// Check if the sum is equal to the original number
	return sum == number
}

func isNumberPerfect (number int) bool {
	if number < 1 {
		return false
	}

	sum := 0
	for i := 1; i < number; i++ {
		if number % i == 0 {
			sum += i
		}
	}

	return sum == number
}

func sumDigits(number int) int {
	sum := 0

	for number != 0 {
		sum += number % 10
		number /= 10
	}

	return sum
} 

func isAlphabet(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func getNumberProperties(number int) []string {
	properties := []string{}
	
	if isNumberArmstrong(number) {
		properties = append(properties, "armstrong")
	}

	if isNumberOdd(number) {
		properties = append(properties, "odd")
	}

	if isNumberEven(number) {
		properties = append(properties, "even")
	}

	return properties
}

func makeNumbersAPIGetRequest(url string) ([]byte, error) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Create a new HTTP client
	client := &http.Client{}

	// Make the GET request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
