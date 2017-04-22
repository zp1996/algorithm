package Algorithm

import (
	"Test"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	v := Test.SortTest(BubbleSort, t)
	if !v {
		t.Errorf("BubbleSort appear error")
	}
}
