package nowcoder

import "strings"

// link  : https://www.nowcoder.com/practice/55fb3c68d08d46119f76ae2df7566880
// title : 验证IP地址
// status: Accepted

func _solveIpV4(ip string) string {
	elems := strings.Split(ip, ".")
	if len(elems) != 4 {
		return "Neither"
	}
	ans := 0
	for idx, size := 0, len(elems); idx < size; idx++ {
		sb := []byte(elems[idx])
		if len(sb) == 0 {
			return "Neither"
		}
		for i, s := 0, len(sb); i < s; i++ {
			if sb[i] == '0' && s != 1 && i == 0 {
				return "Neither"
			}
			ans = ans*10 + int(sb[i]-'0')
		}
		if ans > 255 {
			return "Neither"
		}
		ans = 0
	}
	return "IPv4"
}

func _solveIpV6(ip string) string {
	elems := strings.Split(ip, ":")
	if len(elems) != 8 {
		return "Neither"
	}
	for idx, size := 0, len(elems); idx < size; idx++ {
		sb := []byte(elems[idx])
		if len(sb) == 0 {
			return "Neither"
		}
		for i, s := 0, len(sb); i < s; i++ {
			if sb[i] == '0' && s > 4 {
				return "Neither"
			}
			if (sb[i] >= '0' && sb[i] <= '9') ||
				(sb[i] >= 'A' && sb[i] <= 'F') ||
				(sb[i] >= 'a' && sb[i] <= 'f') {
				continue
			} else {
				return "Neither"
			}
		}
	}
	return "IPv6"
}

func Solve(IP string) string {
	if strings.Contains(IP, ":") {
		return _solveIpV6(IP)
	} else if strings.Contains(IP, ".") {
		return _solveIpV4(IP)
	} else {
		return "Neither"
	}
}
