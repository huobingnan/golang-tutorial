package nowcoder

import (
	"fmt"
	"strings"
	"testing"
)

func TestParseTree(t *testing.T) {
	t.Run("Case1", func(t *testing.T) {
		source := "1,2,3,4,5,6"
		fmt.Println("\tInput : 1,2,3,4,5,6")
		fmt.Println("\tExpect: 1,2,3,4,5,6")
		root := ParseTree(source)
		ret := TreeLevelOrderTraversal(root)
		out := strings.Join(ret, ",")
		fmt.Println("\tOutput:", out)
		if out != source {
			t.Fail()
		}
	})

	t.Run("Case2", func(t *testing.T) {
		input := "1,2,3,#,#,#,4"
		fmt.Println("\tInput :", input)
		fmt.Println("\tExpect:", input)
		root := ParseTree(input)
		output := strings.Join(TreeLevelOrderTraversal(root), ",")
		fmt.Println("\tOutput:", output)
		if output != input {
			t.Fail()
		}
	})

	t.Run("Case3", func(t *testing.T) {
		input := "1,2,3,#,#,#,4,#,#,#,#,#,#,6,7"
		fmt.Println("\tInput :", input)
		fmt.Println("\tExpect:", input)
		root := ParseTree(input)
		output := strings.Join(TreeLevelOrderTraversal(root), ",")
		fmt.Println("\tOutput:", output)
		if input != output {
			t.Fail()
		}
	})

	t.Run("Case4", func(t *testing.T) {
		input := "#,1,2,3,4"
		fmt.Println("\tInput :", input)
		fmt.Println("\tExpect:")
		root := ParseTree(input)
		output := strings.Join(TreeLevelOrderTraversal(root), ",")
		fmt.Println("\tOutput:", output)
		if output != "" {
			t.Fail()
		}
	})
}
