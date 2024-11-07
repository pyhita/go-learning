package main

func Sum(nums []int) int {
	res := 0
	//for i := 0; i < len(nums); i++ {
	//	res += nums[i]
	//}
	for _, num := range nums {
		res += num
	}

	return res
}

func SumAll(numsToSum ...[]int) []int {
	res := make([]int, 0, len(numsToSum))

	for _, nums := range numsToSum {
		res = append(res, Sum(nums))
	}

	return res
}

func SumAllTails(numsToSum ...[]int) []int {
	res := make([]int, 0, len(numsToSum))

	for _, nums := range numsToSum {
		if len(nums) == 0 {
			res = append(res, 0)
		} else {
			res = append(res, Sum(nums[1:]))
		}
	}

	return res
}
