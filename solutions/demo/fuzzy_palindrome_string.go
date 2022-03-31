package solutions

func isLetter(c byte) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}
	if c >= 'A' && c <= 'Z' {
		return true
	}
	return false
}

func FuzzyPalindromeString(s string) bool {
	size := len(s)
	// upper
	chars := []byte(s)
	otherCharCount := 0
	letterCharCount := 0
	table := make([]int, 172)
	for _, char := range chars {
		if isLetter(char) {
			table[char&0xDF] += 1
			letterCharCount += 1
		} else {
			otherCharCount += 1
		}
	}
	sum := 0
	if size%2 != 0 {
		// 如果串是奇数长度的串
		if (otherCharCount+letterCharCount)%2 == 0 {
			return false
		}
		for _, each := range table {
			sum += each
		}
	} else {
		// 如果是偶数长度的串
		if (otherCharCount+letterCharCount)%2 != 0 {
			return false
		}
	}
	return false
}
