package main

// паттерн с двух сторон
func main() {

}

func twoSum(nums []int, target int) []int {
	l := 0
	r := len(nums) - 1 //
	for l < r {
		currSum := nums[l] + nums[r]
		if currSum == target {
			return []int{l, r}
		}
		else if currSum > target {
			r--
		}
		else {
			r++
		}
	}
	return []int{-1, -1}
}

func isPallindrome(s string) bool {
	l:=0
	r:=len(s)-1
	for l < r {
		if s[l] != s[r] {
			continue
		}
		r--;
		l++;
	}
	return true
}

// поиск общих частей массива
func arraysIntersect(a []int, b []int) []int  {
	result :=make([]int, 0)
	p1:=0
	p2:=0
	for p1< len(a) && p2<len(b) {
		if a[p1] < b[p2] {
			p1++
		} else if a[p1] > b[p2] {
			p2++
		} else {
			result = append(result, a[p1])
			p1++
			p2++
		}
	}


	return result
}

func mergeArrays(num1 []int, num2 []int) []int {
	result:= make([]int, 0)
	p1:=0
	p2:=0

	for p1<len(num1) && p2<len(num2) {
		if num1[p1] < num2[p2] {
			result = append(result, num1[p1])
		} else {
			result = append(result, num2[p2])
		}
		p1++
		p2++
	}


	return result
}

// паттерн быстрый и медленный

// переместить все нули в конец
// fast - это ищейка которая смотрит вперед
func moveZeroes(nums []int) []int {
	slow:=0
	fast:= 0
	for fast< len(nums) {
		if nums[fast] != 0  {
			if nums[fast]!=0 {
				nums[slow], nums[fast] = nums[fast], nums[slow]
				slow++ // идет вппедед только после обмена
			}

		}

		fast++ // всегда идет вперед и смотрит
	}

	return nums
}