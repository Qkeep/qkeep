package util

import (
	"log"
	"strconv"
)

func StringToInterface(s string) interface{} {
	return s
}

func InterfaceToString(i interface{}) string {
	return i.(string)
}

func StringToInt(s string, defaultValue int) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("Try to convert %s to int failed, use default value %d", s, defaultValue)
		return defaultValue
	}
	return value
}

func SnakeCaseToUpperCamelCase(s string) string {
	var result string
	for _, word := range SplitBy(s, "_") {
		result += UpperFirst(word)
	}
	return result
}

func UpperFirst(s string) string {
	if len(s) == 0 {
		return ""
	}
	return string(s[0]-32) + s[1:]
}

func SplitBy(s string, sep string) []string {
	var result []string
	var word string
	for _, c := range s {
		if string(c) == sep {
			result = append(result, word)
			word = ""
		} else {
			word += string(c)
		}
	}
	result = append(result, word)
	return result
}
