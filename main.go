package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

import _ "github.com/heroku/x/hmetrics/onload"

func isPalindrome(s string) bool {
	s = strings.ToLower(strings.TrimSpace(s))
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func rot13(s string) string {
	var result strings.Builder
	result.Grow(len(s))
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z':
			result.WriteRune('a' + (r-'a'+13)%26)
		case r >= 'A' && r <= 'Z':
			result.WriteRune('A' + (r-'A'+13)%26)
		default:
			result.WriteRune(r)
		}
	}
	return result.String()
}

// not really random or performant, but good enough for this purpose
func randomBool() bool {
	return rand.Intn(2) == 1
}

func spongebob(s string) string {
	var result strings.Builder
	upper := randomBool()
	for _, r := range s {
		if upper {
			result.WriteString(strings.ToUpper(string(r)))
		} else {
			result.WriteString(strings.ToLower(string(r)))
		}
		upper = !upper
	}
	return result.String()
}

type requestPayload struct {
	Input *string
}

func extractPayload(r *http.Request) (requestPayload, error) {
	decoder := json.NewDecoder(r.Body)
	var t requestPayload
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	return t, err
}

// palindromeHandler checks if the provided string is a palindrome.
func palindromeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := extractPayload(r)

	result := map[string]interface{}{
		"original":     *t.Input,
		"isPalindrome": isPalindrome(*t.Input),
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		return
	}
}

// rot13Handler applies ROT13 encoding to the provided string.
func rot13Handler(w http.ResponseWriter, r *http.Request) {
	t, err := extractPayload(r)

	result := map[string]string{
		"original": *t.Input,
		"rot13":    rot13(*t.Input),
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		return
	}
}

// spongebobHandler applies spongebob encoding to the provided string.
func spongebobHandler(w http.ResponseWriter, r *http.Request) {
	t, err := extractPayload(r)

	result := map[string]string{
		"original":  *t.Input,
		"spongebob": spongebob(*t.Input),
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		return
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	result := map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Service is healthy",
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/is-palindrome", palindromeHandler)
	http.HandleFunc("/rot13", rot13Handler)
	http.HandleFunc("/spongebob", spongebobHandler)
	http.HandleFunc("/health", healthHandler)

	appPort := ":" + os.Getenv("PORT")
	if appPort == ":" {
		appPort = ":8080"
	}
	err := http.ListenAndServe(appPort, nil)
	if err != nil {
		return
	}
}
