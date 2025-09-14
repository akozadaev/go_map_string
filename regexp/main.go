package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := `Многострочный
текст`
	pattern := "текbст"
	r := regexp.MustCompile(pattern)
	if r.MatchString(text) {
		fmt.Println("Найдено вхождение")
	} else {
		fmt.Println("Не найдено вхождение")
	}

	//regexp.Compile() //
	//r.FindString()   //Первое совпадение
	//r.FindStringSubmatch(text)
	//r.Re(text)

	//fmt.Println(text)
}
