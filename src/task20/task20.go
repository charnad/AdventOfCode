package task20

import (
	"fmt"
	"math"
)

func sliceSum(ints []int) int {
	result := 0
	for _, i := range ints {
		result += i
	}
	return result
}

func divisors(num int) []int {
	result := []int{}

	for i := 1; i <= int(math.Sqrt(float64(num))); i++ {
		if num % i == 0 {
			result = append(result, i)
		}

		if int(num / i) != i && num % int(num / i) == 0 {
			result = append(result, int(num / i))
		}
	}

	return result
}

func divisorsSum(num int) int {
	return sliceSum(divisors(num))
}

func min(nums []int) int {
	var result int = 0

	for _, i := range nums {
		if i < result {
			result = i
		}
	}

	return result
}

func max(nums []int) int {
	var result int = 0

	for _, i := range nums {
		if i > result {
			result = i
		}
	}

	return result
}

func Solve() {
//	targetPresents := 34000000
//	maxSum := targetPresents / 10
//
//	i := 1
//
//	fmt.Println(divisors(786240))
//
//	return
//
//	// The smallest number which has a sum of divisors greater or equal to X
//	for ; i < maxSum; i++ {
//		x := divisorsSum(i)
//		if x >= maxSum {
//			fmt.Println("Result:", i)
//			break
//		}
//	}

	for i := 3400001 ;; i += 11 {
		divs := divisors(i)
		fmt.Println(divs)

	}
}
