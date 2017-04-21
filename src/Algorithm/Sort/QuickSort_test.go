package Algorithm

import (
	"Test"
	"testing"
)

func TestQucikSort(t *testing.T) {
	v := Test.SortTest(QucikSort)
	if !v {
		t.Errorf("QuickSort is error")
	}
}
