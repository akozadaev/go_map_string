package main

import (
	"fmt"
	"regexp"
)

func main() {
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailPattern)
	testEmails := []string{
		"user@example.com",
		"test.email+tag@domain.co.uk",
		"invalid.email",
		"@noat.com",
		"noat.",
	}
	for _, email := range testEmails {
		if re.MatchString(email) {
			fmt.Printf("%s — валидный email\n", email)
		} else {
			fmt.Printf("%s — невалидный email\n", email)
		}
	}
}
