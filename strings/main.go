package main

import (
	"fmt"
	"strings"
)

func main() {
	//strEng := "q"
	//strRu := "й"
	//strY := "🎶"
	//strI := "世"
	////strI := "世界"
	//fmt.Println(len(strEng))
	//fmt.Println(len(strRu))
	//fmt.Println(len(strI))
	//fmt.Println(len(strY))
	//IterateByte("моя строка")
	//IterateStr("моя строка")
	//StringsPackage()
	StrBuilder()
}

func IterateByte(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
		fmt.Printf("%c", s[i])
		fmt.Println(string(s[i]))
	}
}
func IterateStr(s string) {
	for _, r := range s {
		fmt.Printf("%c\n", r)
	}
}

func StringsPackage() {
	fmt.Println(strings.Contains("Golang", "Go"))
	fmt.Println(strings.Replace("Golang", "lang", "", -1))
	fmt.Println(strings.Replace("GolangGo", "Go", "", -1))
	fmt.Println(strings.Replace("GolangGo", "Go", "", 1))
	fmt.Println(strings.Replace("GolangGo", "Go", "", 2))
}

func StrBuilder() {
	var builder strings.Builder
	builder.WriteString("Golang, ")
	builder.WriteString("Go, ")
	builder.WriteString("Go lang")
	fmt.Println(builder.String())
}
