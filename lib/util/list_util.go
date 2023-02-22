package util

func FilterStringArr(arr []string, filter func(string) bool) []string {
	var result []string
	for _, v := range arr {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

func MapStringArr(arr []string, mapper func(string) string) []string {
	var result []string
	for _, v := range arr {
		result = append(result, mapper(v))
	}
	return result
}
