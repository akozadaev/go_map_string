package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{
		"score": 30,
		"age":   30,
		"r":     30,
		"z":     30,
	}
	for i := 0; i < 100; i++ {
		fmt.Println("===")
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
	fmt.Println("===+++++===")
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	//sort.Reverse(sort.Strings(keys))
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, m[key])

	}
	//
	//fmt.Println(m)
	//fmt.Println(m["age"])
	//m["age"] = 23
	//
	//fmt.Println(m["age"])
	//delete(m, "age")
	//fmt.Println(m["age"])
	////_ := m["age"]
	//fmt.Println(m)
	//m["age"] = 1
	//
	//val, ok := m["age"]
	//if ok {
	//	fmt.Println("str1 ", val)
	//} else {
	//	fmt.Println("no value ")
	//}

	//fmt.Println(m["age"])
	//fmt.Println(m["age2"])
	//fmt.Println("=== ADD ===")
	//fmt.Println(m)
	//m["age2"] = 1
	//fmt.Println(m)
	//
	//var m1 map[string]int // nil map
	//fmt.Println(m1)
	//

	var m1 map[string]int // nil map

	fmt.Println(len(m1))
	m2 := make(map[string]int)
	fmt.Println(len(m2))
	m3 := make(map[string]int, 10)
	fmt.Println(m3)

	//s := "gopher"
	//fmt.Printf("Hello and welcome, %s!\n", s)
	//
	//for i := 1; i <= 5; i++ {
	//	fmt.Println("i =", 100/i)
	//}
}
