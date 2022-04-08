package utils

func ClearEmptyString(arr []string) []string {
	result := make([]string, 0)
	for _, s := range arr {
		if s != "" {
			result = append(result, s)
		}
	}
	return result
}
