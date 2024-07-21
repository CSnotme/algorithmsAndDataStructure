package leetcode

import (
	"fmt"
	"testing"
)

func Test_946(t *testing.T) {

	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}
	ret := validateStackSequences(pushed, popped)
	fmt.Println(ret)
}
