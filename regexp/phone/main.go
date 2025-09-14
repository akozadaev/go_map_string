package main

import (
	"fmt"
	"regexp"
)

func main() {
	// +7 (111) 111-11-11
	phonePattern := `\+7\s*\(\d{3}\)\s*\d{3}-\d{2}-\d{2}`
	re := regexp.MustCompile(phonePattern)
	test := ` +7 (111) 111-11-11 +7 (111)    111-44-22`

	match := re.FindAllString(test, -1)
	fmt.Println("Найдено: ")

	for _, m := range match {
		fmt.Println(m)
	}
}
