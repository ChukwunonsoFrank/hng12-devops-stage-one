package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
	"os"
	"github.com/joho/godotenv"
)

type Response struct {
	Number     int                 `json:"number"`
	IsPrime    bool                `json:"is_prime"`
	IsPerfect  bool                `json:"is_perfect"`
	Properties []string            `json:"properties"`
	DigitSum   int                 `json:"digit_sum"`
	FunFact    string              `json:"fun_fact"`
}

type Error struct {
	Number     string                 `json:"number"`
	Error      bool                   `json:"error"`
}

// CORS middleware
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Adjust for specific domains if needed
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			return // Return immediately for preflight requests
		}
		next.ServeHTTP(w, r)
	})
}

// Basic API handler
func classifyNumberHandler(w http.ResponseWriter, r *http.Request) {
	var response interface{}
	w.Header().Set("Content-Type", "application/json")
	
	// Parse query parameters
	number := r.URL.Query().Get("number")

	// Validation checks
	if number == "" {
        response = Error{
			Number: "alphabet",
			Error:  true,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
    }

	if isAlphabet(number) {
		response = Error{
			Number: "alphabet",
			Error:  true,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	numberConvertedToInteger, numberConvertedToIntegerErr := strconv.Atoi(number)

	if numberConvertedToIntegerErr != nil {
		response = Error{
			Number: "alphabet",
			Error:  true,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	numbersAPIResponseBody, numbersAPIResponseErr := makeNumbersAPIGetRequest("http://numbersapi.com/" + number + "/math")
	if numbersAPIResponseErr != nil {
		log.Println(numbersAPIResponseErr)
		return
	}

	response = Response{
		Number:     numberConvertedToInteger,
		IsPrime:    isNumberPrime(numberConvertedToInteger),
		IsPerfect:  isNumberPerfect(numberConvertedToInteger),
		Properties: getNumberProperties(numberConvertedToInteger),
		DigitSum:   sumDigits(numberConvertedToInteger),
		FunFact:    string(numbersAPIResponseBody),
	}

	json.NewEncoder(w).Encode(response)
}

// Main function
func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not found in the environment!")
		return
	}
	
	mux := http.NewServeMux()

	// Define routes
	mux.HandleFunc("/api/classify-number", classifyNumberHandler)

	// Start the server with CORS enabled
	server := &http.Server{
		Addr:         "0.0.0.0:" + portString,
		Handler:      enableCORS(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Starting server on port " + portString + "..." )
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}