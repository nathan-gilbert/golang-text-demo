package textulizer

import (
	"encoding/json"
	"net/http"
	"strings"
)

// isPalindrome checks if a string is a palindrome.
func isPalindrome(s string) bool {
	s = strings.ToLower(strings.TrimSpace(s))
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// rot13 applies ROT13 encoding to a string.
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

// palindromeHandler checks if the provided string is a palindrome.
func palindromeHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	input := query.Get("input")

	result := map[string]interface{}{
		"input":      input,
		"palindrome": isPalindrome(input),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// rot13Handler applies ROT13 encoding to the provided string.
func rot13Handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	input := query.Get("input")

	result := map[string]string{
		"input": input,
		"rot13": rot13(input),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// spongebobHandler applies spongebob encoding to the provided string.

func main() {
	http.HandleFunc("/is-palindrome", palindromeHandler)
	http.HandleFunc("/rot13", rot13Handler)

	http.ListenAndServe(":8080", nil)
}
