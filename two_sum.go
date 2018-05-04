package main

func twoSum(nums []int, target int) []int {
	var result []int
	length := len(nums)
	for i, num := range nums {
		for j := i + 1; j < length; j++ {
			if (num + nums[j]) == target {
				result = append(result, i, j)
			}
		}
	}
	return result
}

func main() {
	nums := []int{2, 7, 11, 15}
	twoSum(nums, 9)
}
