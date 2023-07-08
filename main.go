package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(int64(42.0))
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func MinInt(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func IsValid(id int, text string) bool {
	return id > 0 && text != ""
}

func Greetings(name string) string {
	name = strings.Trim(name, " ")
	name = strings.ToLower(name)
	name = strings.Title(name)

	return fmt.Sprintf("Привет, %s!", name)
}

func DomainForLocale(domain string, locale string) string {
	if locale != "" {
		return fmt.Sprintf("%s.%s", locale, domain)
	}

	return fmt.Sprintf("en.%s", domain)
}

func ModifySpaces(s, mode string) string {
	switch mode {
		case "dash": 
			return strings.ReplaceAll(s, " ", "-")
		case "underscore": 
			return strings.ReplaceAll(s, " ", "_")
		default: 
			return strings.ReplaceAll(s, " ", "*")	
	}
}

type UserCreateRequest struct {
	FirstName string
	Age       int
}

func Validate(req UserCreateRequest) string {
	message := "invalid request"

	// валидация имени
	if (req.FirstName == "" || strings.Contains(req.FirstName, " ")) {
		return message
	}

	// валидация возраста
	if (req.Age <= 0 || req.Age > 150) {
		return message
	}

	return ""
}