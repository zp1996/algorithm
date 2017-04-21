package Test

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

const testLen int = 100

func getSlice() []map[string][]int {
	sliceMap := make([]map[string][]int, testLen)
	for i := 0; i < testLen; i++ {
		uniqueMap := make(map[string][]int)
		rand.Seed(int64(time.Now().Nanosecond()))
		len := rand.Intn(1000)
		slice := make([]int, len)
		for j := 0; j < len; j++ {
			slice[j] = rand.Intn(10000)
		}
		sorted := make([]int, len, cap(slice))
		copy(sorted, slice)
		sort.Ints(sorted)
		uniqueMap["init"] = slice
		uniqueMap["sorted"] = sorted
		sliceMap[i] = uniqueMap
	}
	return sliceMap
}

func equal(init, sorted []int) bool {
	if len(init) != len(sorted) {
		return false
	}
	for i, v := range init {
		if v != sorted[i] {
			return false
		}
	}
	return true
}

func SortTest(fn func([]int) []int, t *testing.T) bool {
	sliceMap := getSlice()
	for _, v := range sliceMap {
		init, initStatus := v["init"]
		sorted, sortedStatus := v["sorted"]
		attention := make([]int, len(init))
		copy(attention, init)
		if !(initStatus && sortedStatus && equal(fn(init), sorted)) {
			t.Errorf("%v\n对上面数组进行排序时出现错误", attention)
			return false
		}
	}
	return true
}
