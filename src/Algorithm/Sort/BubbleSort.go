package Algorithm

func BubbleSort(arr []int) []int {
	len := len(arr)
	flag := false
	for i := 0; i < len; i++ {
		for j := len - 1; j > i; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return arr
}
