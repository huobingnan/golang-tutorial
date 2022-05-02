package nowcoder

import "math"

// link  : https://www.nowcoder.com/practice/11ae12e8c6fe48f883cad618c2e81475
// title : 大数加法
// status: Accepted

func SolveBm86(s, t string) string {
	carry, p1, p2, e1, e2 := byte(0), len(s)-1, len(t)-1, -1, -1
	ret := make([]byte, 0, int(math.Max(float64(e1), float64(e2)))+1)
	for {
		if p1 == e1 && p2 == e2 {
			break
		} else if p1 != e1 && p2 != e2 {
			bit := (s[p1] - '0') + (t[p2] - '0') + carry
			ret = append(ret, (bit%10)+'0')
			if bit >= 10 {
				carry = 1
			} else {
				carry = 0
			}
			p1, p2 = p1-1, p2-1
		} else if p1 != e1 {
			bit := (s[p1] - '0') + carry
			ret = append(ret, bit%10+'0')
			if bit >= 10 {
				carry = 1
			} else {
				carry = 0
			}
			p1--
		} else {
			bit := (t[p2] - '0') + carry
			ret = append(ret, bit%10+'0')
			if bit >= 10 {
				carry = 1
			} else {
				carry = 0
			}
			p2--
		}
	}
	if carry != 0 {
		ret = append(ret, carry+'0')
		carry = 0
	}
	for l, r := 0, len(ret)-1; l < r; l, r = l+1, r-1 {
		ret[l], ret[r] = ret[r], ret[l]
	}
	return string(ret)
}
