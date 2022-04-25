package leetcode

import (
	"strconv"
	"strings"
)

// link: https://leetcode-cn.com/problems/restore-ip-addresses/
// Accepted

func _doSearch(seq []byte, start int, itemCount int, temp []string, collector *[]string) {
	if start >= len(seq) {
		if itemCount == 4 {
			// 收集结果
			*collector = append(*collector, strings.Join(temp, "."))
		}
		return
	}
	if itemCount >= 4 {
		return
	}
	ans := 0
	// 连续的三个序列
	if start+3 <= len(seq) && seq[start] != '0' {
		for idx := start; idx < start+3; idx++ {
			ans = ans*10 + int(seq[idx]-'0')
		}
		if ans <= 255 {
			temp[itemCount] = strconv.Itoa(ans)
			_doSearch(seq, start+3, itemCount+1, temp, collector)
		}
	}
	// 连续两个序列
	if start+2 <= len(seq) && seq[start] != '0' {
		if ans == 0 {
			for idx := start; idx < start+2; idx++ {
				ans = ans*10 + int(seq[idx]-'0')
			}
		} else {
			ans = ans / 10
		}
		temp[itemCount] = strconv.Itoa(ans)
		_doSearch(seq, start+2, itemCount+1, temp, collector)
	}
	// 连续一个序列
	if start+1 <= len(seq) {
		if ans == 0 {
			ans = ans*10 + int(seq[start]-'0')
		} else {
			ans = ans / 10
		}
		temp[itemCount] = strconv.Itoa(ans)
		_doSearch(seq, start+1, itemCount+1, temp, collector)
	}
}

func RestoreIpAddresses(s string) []string {
	sb := []byte(s)
	collector := make([]string, 0, 10)
	_doSearch(sb, 0, 0, make([]string, 4), &collector)
	return collector
}
