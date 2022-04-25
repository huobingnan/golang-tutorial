package demo

import "testing"

func TestValidIPAddress(t *testing.T) {
	t.Run("Case1", func(t *testing.T) {
		ret := ValidIPAddress("20522518135")
		for _, each := range ret {
			t.Log(each)
		}
	})
	t.Run("Case2", func(t *testing.T) {
		ret := ValidIPAddress("12121203")
		for _, each := range ret {
			t.Log(each)
		}
	})
}
