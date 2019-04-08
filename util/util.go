package util

func TruncateV2(input []byte) []byte {
	var i int
	// Loop through the array in reverse until the first non-zero byte is found
	for i = len(input); i > 0; i-- {
		if input[i-1] != 0 {
			return input[:i]
		}
	}

	return input
}

// CStrLen returns the length of the null-terminated C string contained in the input buffer
func CStrLen(input []byte) int {
	for i := 0; i < len(n); i++ {
		if n[i] == 0 {
			return i
		}
	}
	return len(n)
}
