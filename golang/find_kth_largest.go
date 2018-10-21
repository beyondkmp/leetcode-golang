package main

import "fmt"

func less(a, b int) bool {
	if a < b {
		return true
	}
	return false
}

func partition(a []int, l, r int) int {
	i := l + 1
	j := r
	v := a[l]

	for {
		for ; i <= r && less(a[i], v); i++ {
		}

		for ; j >= l && less(v, a[j]); j-- {
		}

		if i >= j {
			break
		}

		a[i], a[j] = a[j], a[i]
		i++
		j--
	}

	a[j], a[l] = a[l], a[j]
	return j
}

func findKthLargest(nums []int, k int) int {
	p := len(nums) - k
	l := 0
	r := len(nums) - 1

	for l <= r {
		i := partition(nums, l, r)
		switch {
		case i == p:
			return nums[i]
		case i < p:
			l = i + 1
		case i > p:
			r = i - 1
		}
	}
	return nums[l]
}

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(nums, 4))
}
