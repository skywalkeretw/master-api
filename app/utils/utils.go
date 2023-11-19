package utils

// Contains check to see if a string is contained in a slice of string
func Contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
