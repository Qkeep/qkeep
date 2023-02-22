package util

func StringToInterface(s string) interface{} {
	return s
}

func InterfaceToString(i interface{}) string {
	return i.(string)
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

	// Check if the first character is already uppercase or number
	if (s[0] >= 65 && s[0] <= 90) || (s[0] >= 48 && s[0] <= 57) {
		return s
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
