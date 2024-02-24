package util

func ParseBool(str string) bool {
	switch str {
	case "1", "t", "T", "true", "TRUE", "True", "Y", "y", "yes", "YES", "Yes":
		return true
	case "0", "f", "F", "false", "FALSE", "False", "n", "N", "no", "NO", "No":
		return false
	}
	return false
}
