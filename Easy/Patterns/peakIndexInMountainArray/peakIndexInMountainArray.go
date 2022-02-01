package peakIndexInMountainArray

import "fmt"

func Example() {
	fmt.Println("\nPeak Index in a Mountain Array")
	nums := []int{1, 57, 58, 74, 88, 93, 98, 97, 96, 91, 90, 78, 77, 74, 71, 68, 61, 50, 42, 38, 35, 34, 26, 20, 15, 14, 5, 4, 2}
	fmt.Println("nums =", nums)
	result := peakIndexInMountainArray(nums)
	fmt.Println("result =", result)
}

func peakIndexInMountainArray(arr []int) int {
	//p - второй элемент
	p := 1
	//проверочка
	if check(arr, p) {
		return p
	}
	//p - средний элемент
	p = len(arr) / 2
	//t - сдвиг
	t := p / 2
	if check(arr, p) {
		return p
	}
	for !check(arr, p) && t > 1 {
		if arr[p] > arr[p+1] {
			p -= t
			t = t / 2
			continue
		} else {
			p += t
			t = t / 2
			continue
		}
	}
	for !check(arr, p) {
		if arr[p] > arr[p+1] {
			p -= 1
		} else {
			p += t
		}
	}
	return p
}

func check(arr []int, p int) bool {
	return arr[p-1] < arr[p] && arr[p+1] < arr[p]
}
