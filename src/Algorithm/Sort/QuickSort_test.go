package Algorithm

import (
	"Test"
	"testing"
)

func TestQucikSort(t *testing.T) {
	v := Test.SortTest(QucikSort, t)
	if !v {
		t.Errorf("QuickSort appear error")
	}
}
