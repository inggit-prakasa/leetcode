package main

// func main() {
// 	fmt.Println(findXSum([]int{1, 2, 3, 4, 5}, 3, 2)) // Example usage
// }

// func findXSum(nums []int, k int, x int) []int {
// 	resp := make([]int, 0)

// 	sort.Ints(nums)

// 	windowSum := 0
// 	for windowSum < len(nums)-k+1 {
// 		semt := make([]int, 0)
// 		m := make(map[int]int)

// 		for i := windowSum; i < k+windowSum; i++ {
// 			m[nums[i]]++

// 			semt = append(semt, nums[i])
// 		}

// 		unique := []int{}
// 		for num := range freq {
// 			unique = append(unique, num)
// 		}

// 		windowSum++
// 	}

// 	return resp
// }
