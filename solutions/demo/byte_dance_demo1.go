package demo

import (
	"container/list"
	"strconv"
	"strings"
)

func _doSearch(seq []byte, start int, itemCount int, temp *list.List, collector *[]string) {
	if itemCount > 4 {
		return
	}
	if start >= len(seq) {
		if itemCount == 4 {
			// 收集结果
			for e := temp.Front(); e != nil; e = e.Next() {
				v := e.Value.(int)
				*collector = append(*collector, strconv.Itoa(v))
			}
		}
		return
	}
	ans := 0
	// 连续的三个序列
	if start+3 <= len(seq) {
		for idx := start; idx < start+3; idx++ {
			ans = ans*10 + int(seq[idx]-'0')
		}
		if ans <= 255 {
			temp.PushBack(ans)
			_doSearch(seq, start+3, itemCount+1, temp, collector)
			temp.Remove(temp.Back())
		}
		if ans/10 <= 255 {
			temp.PushBack(ans / 10)
			_doSearch(seq, start+2, itemCount+1, temp, collector)
			temp.Remove(temp.Back())
		}
		if ans/100 <= 255 {
			temp.PushBack(ans / 100)
			_doSearch(seq, start+1, itemCount+1, temp, collector)
			temp.Remove(temp.Back())
		}
	} else if start+2 <= len(seq) {
		for idx := start; idx < start+2; idx++ {
			ans = ans*10 + int(seq[idx]-'0')
		}
		if ans <= 255 {
			temp.PushBack(ans)
			_doSearch(seq, start+2, itemCount+1, temp, collector)
			temp.Remove(temp.Back())
		}
		if ans/10 <= 255 {
			temp.PushBack(ans / 10)
			_doSearch(seq, start+1, itemCount+1, temp, collector)
			temp.Remove(temp.Back())
		}
	} else if start+1 <= len(seq) {
		ans = ans*10 + int(seq[start]-'0')
		if ans <= 255 {
			temp.PushBack(ans)
			_doSearch(seq, start+1, itemCount+1, temp, collector)
			temp.Remove(temp.Back())
		}
	}
}

func ValidIPAddress(input string) []string {
	ib := []byte(input)
	tempRes := list.New()
	collector := make([]string, 0, 20)
	_doSearch(ib, 0, 0, tempRes, &collector)
	ts := make([]string, 4)
	ti := 0
	ret := make([]string, 0, 5)
	for _, each := range collector {
		ts[ti] = each
		ti++
		if ti%4 == 0 {
			ret = append(ret, strings.Join(ts, "."))
			ti = 0
		}
	}
	return ret
}
