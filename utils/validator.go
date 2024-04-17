package utils

import "regexp"

// IsValidEmail memeriksa apakah email memiliki format yang valid
func IsValidEmail(email string) bool {
	// Pola regular expression untuk validasi email
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Validasi email dengan pola regular expression
	match, _ := regexp.MatchString(emailRegex, email)
	return match
}
