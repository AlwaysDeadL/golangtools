package strings

// HasPrefix is a method to check if contains prefix str
// 判断字符串是否包含前缀
func HasPrefix(s, prefix string) bool {
	return len(s) > len(prefix) && s[:len(prefix)] == prefix
}

// HasSuffix is a method to check if contains the suffix str
func HasSuffix(s, suffix string) bool {
	return len(s) > len(suffix) && s[len(s)-len(suffix):] == suffix
}

// Contains is a method to check if contains the substr
func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
