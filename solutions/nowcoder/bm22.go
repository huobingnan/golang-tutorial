package nowcoder

// link: https://www.nowcoder.com/practice/2b317e02f14247a49ffdbdba315459e7
// Accepted

func _parseModifyCode(version []byte) (int, int) {
	ans := 0
	for idx, n := range version {
		if n == '.' {
			return ans, idx + 1
		}
		ans = ans*10 + int(n-'0')
	}
	return ans, len(version)
}

func Compare(version1, version2 string) int {
	vb1 := []byte(version1)
	vb2 := []byte(version2)
	for {
		v1, n1 := _parseModifyCode(vb1)
		v2, n2 := _parseModifyCode(vb2)
		// compare
		if v1 > v2 {
			return 1
		} else if v1 < v2 {
			return -1
		}
		if n1 == len(vb1) && n2 == len(vb2) {
			return 0
		}
		if n1 <= len(vb1) {
			vb1 = vb1[n1:]
		}
		if n2 <= len(vb2) {
			vb2 = vb2[n2:]
		}
	}
}
