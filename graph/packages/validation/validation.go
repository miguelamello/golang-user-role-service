package validation

import (
	"regexp"
	"strings"
	"github.com/google/uuid"
)

func VerifyUUID(uuidString string) bool {
	_, err := uuid.Parse(uuidString)
	if err != nil {
		return false
	} else {
		return true
	}
}

// Verify if the email string is valid
func isEmail(email string) bool {
	// Define a regular expression for validating email addresses
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	// Check if the email address matches the regular expression
	if re.MatchString(email) {
		return true
	} else {
		return false
	}
}

// Verify if the email string exists and is valid
func ValidateEmailString(email string) bool {
	cEmail := strings.Trim(email, " ")
	if len(cEmail) == 0 { 
		return false 
	} else { 
		if isEmail(cEmail) {
			return true 
		} else {			
			return false
		}
	}
}

// Verify if the bearer string is valid
func ValidateBearerString(bearer string) bool {
	cBearer := strings.Trim(bearer, " ")
	if len(cBearer) == 0 { 
		return false 
	} else { 
		return true 
	}
}
