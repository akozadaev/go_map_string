package main

import (
	"fmt"
	"regexp"
)

func main() {
	urlPattern := `^(https?://)([a-zA-Z0-9.-]+)(:[0-9]+)?(/.*)?$` // ЗДЕСЬ круглые скобки - это группы!
	re := regexp.MustCompile(urlPattern)
	urls := []string{
		"https://google.com:8080/path/to/page",
		"https://google.com/path/to/page",
		"http://yandex.ru",
		"http://ya.ru",
		"http://ya.ru:443",
		"https://example.org",
	}
	for _, url := range urls {
		matches := re.FindStringSubmatch(url) // возвращает полное совпаден
		if len(matches) > 0 {
			fmt.Printf("URL: %s\n", url)
			fmt.Printf("Протокол: %s\n", matches[1]) // (https?://)
			fmt.Printf("Домен: %s\n", matches[2])    // ([a-zA-Z0-9.-]+)
			fmt.Printf("Порт: %s\n", matches[3])     // (:[0-9]+)?
			fmt.Printf("Путь: %s\n", matches[4])     // (/.*)
			fmt.Println()
		}
	}
}
