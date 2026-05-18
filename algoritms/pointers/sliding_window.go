package main

import "strconv"

// окна фикс длины
func maxSum(nums []int, k int) int {
	windowSum := 0

	for i := 0; i < k; i++ {
		windowSum += nums[i] // собрали сумму в окне

	}
	res := windowSum
	for right := k; right < len(nums); right++ { //  движение окна
		l := right - k                                 // выставили границу окна
		windowSum := windowSum + nums[right] - nums[l] // отняли и прибавили эл-т
		if windowSum > res {
			res = windowSum
		}
	}

	return res
}

func maxMul(nums []int, k int) int {
	windowMul := 0
	res := 0
	for i := 0; i < k; i++ {
		res = res * nums[i]
	}
	res = windowMul
	for right := k; right < len(nums); right++ {
		l := right - k
		windowMul = (windowMul * nums[right]) / nums[l]
	}
	return res
}

// непересекающиеся окна
func intArr(nums []int) []string {
	l := 0
	r := 0
	res := make([]string, 0)

	for l < len(nums) {
		for r+1 < len(nums) && nums[r]+1 == nums[r+1] {
			r++
		}
		if r != l {
			res = append(res, strconv.Itoa(l)+"->"+strconv.Itoa(r))
		} else {
			res = append(res, strconv.Itoa(l))
		}
		l = r + 1
		r = r + 1
	}

}
