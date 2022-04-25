package leetcode

import "testing"

func TestRestoreIpAddresses(t *testing.T) {
	t.Run("25525511135", func(t *testing.T) {
		ret := RestoreIpAddresses("25525511135")
		for _, each := range ret {
			t.Log(each)
		}
	})
	t.Run("0000", func(t *testing.T) {
		ret := RestoreIpAddresses("0000")
		for _, each := range ret {
			t.Log(each)
		}
	})
	t.Run("101023", func(t *testing.T) {
		for _, each := range RestoreIpAddresses("101023") {
			t.Log(each)
		}
	})
}
