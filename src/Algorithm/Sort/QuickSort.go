package Algorithm

func qSort(arr []int, left, right int) {
	current := arr[left]
	i := left
	j := right
	for i < j {
		for arr[j] > current && i < j {
			j--
		}
		for arr[i] <= current && i < j {
			i++
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[left] = arr[i]
	arr[i] = current
	if left < j-1 {
		qSort(arr, left, j-1)
	}
	if right > j+1 {
		qSort(arr, j+1, right)
	}
}

func QucikSort(arr []int) []int {
	qSort(arr, 0, len(arr)-1)
	return arr
}
