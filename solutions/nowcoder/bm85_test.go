package nowcoder

import "testing"

func TestSolve(t *testing.T) {
	t.Run("172.16.254.1", func(t *testing.T) {
		ip := "172.16.254.1"
		if Solve(ip) != "IPv4" {
			t.Fail()
		}
	})

	t.Run("256.256.256.256", func(t *testing.T) {
		ip := "256.256.256.256"
		if Solve(ip) != "Neither" {
			t.Fail()
		}
	})

	t.Run("0.121.0.232", func(t *testing.T) {
		ip := "0.121.0.232"
		if Solve(ip) != "IPv4" {
			t.Fail()
		}
	})

	t.Run("121.122.012.1", func(t *testing.T) {
		ip := "121.122.012.1"
		if Solve(ip) != "Neither" {
			t.Fail()
		}
	})

	t.Run("2001:0db8:85a3:0:0:8A2E:0370:7334", func(t *testing.T) {
		ip := "2001:0db8:85a3:0:0:8A2E:0370:7334"
		if Solve(ip) != "IPv6" {
			t.Fail()
		}
	})

	t.Run("02001:0db8:85a3:0000:0000:8a2e:0370:7334", func(t *testing.T) {
		ip := "02001:0db8:85a3:0000:0000:8a2e:0370:7334"
		if Solve(ip) != "Neither" {
			t.Fail()
		}
	})
	t.Run("2001:0db8:85a3::8A2E:0370:7334", func(t *testing.T) {
		ip := "2001:0db8:85a3::8A2E:0370:7334"
		if Solve(ip) != "Neither" {
			t.Fail()
		}
	})
}
