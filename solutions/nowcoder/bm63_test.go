package nowcoder

import "testing"

func TestJumpFloor(t *testing.T) {
	t.Run("Case1", func(t *testing.T) {
		if JumpFloor(2) != 2 {
			t.Fail()
		}
	})

	t.Run("Case2", func(t *testing.T) {
		if JumpFloor(7) != 21 {
			t.Fail()
		}
	})
}
