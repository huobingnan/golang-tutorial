package nowcoder

import "testing"

func TestCompare(t *testing.T) {
	t.Run("Case1", func(t *testing.T) {
		r := Compare("012.0.0", "2.1.1")
		if r != 1 {
			t.Fail()
		}
	})
	t.Run("Case2", func(t *testing.T) {
		r := Compare("1.1", "2.1")
		if r != -1 {
			t.Fail()
		}
	})
	t.Run("Case3", func(t *testing.T) {
		if Compare("1.1", "1.01") != 0 {
			t.Fail()
		}
	})
	t.Run("Case4", func(t *testing.T) {
		r := Compare("2.0.1", "2")
		if r != 1 {
			t.Log("Expect 1", "Actual", r)
			t.Fail()
		}
	})

	t.Run("Case5", func(t *testing.T) {
		r := Compare("0.226", "0.36")
		if r != 1 {
			t.Log("Expect 1", "Actual", r)
			t.Fail()
		}
	})
}
