package nowcoder

import "testing"

func TestSolveBm86(t *testing.T) {
	t.Run("1+99=100", func(t *testing.T) {
		ret := SolveBm86("1", "99")
		t.Log(ret)
		if ret != "100" {
			t.Fail()
		}
	})
	t.Run("114514+''=114514", func(t *testing.T) {
		ret := SolveBm86("114514", "")
		t.Log(ret)
		if ret != "114514" {
			t.Fail()
		}
	})
	t.Run("100+0=100", func(t *testing.T) {
		ret := SolveBm86("100", "0")
		t.Log(ret)
		if ret != "100" {
			t.Fail()
		}
	})
}
