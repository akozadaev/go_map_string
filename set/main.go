package main

import "fmt"

type Set map[string]struct{}

func main() {

	set := make(Set)
	set.Add("str1")
	set.Add("str2")
	set.Add("str2")
	fmt.Println(set)
	b := set.Contains("str1")
	fmt.Println(b)
	b = set.Contains("str2")
	fmt.Println(b)
	b = set.Contains("str3")
	fmt.Println(b)
	set.Remove("str2")
	b = set.Contains("str2")
	fmt.Println(b)
	fmt.Println(set)
}

// Добавление
func (s Set) Add(value string) {
	s[value] = struct{}{}
}

// Удаление
func (s Set) Remove(value string) {
	delete(s, value)
}

// Проверка наличия
func (s Set) Contains(value string) bool {
	_, ok := s[value]
	return ok
}

// Получение всех элементов
func (s Set) Values() []string {
	values := make([]string, 0, len(s))
	for value := range s {
		values = append(values, value)
	}
	return values
}

// Размер
func (s Set) Size() int {
	return len(s)

}
