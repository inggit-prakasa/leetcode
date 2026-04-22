package main

func sumFourDivisors(nums []int) int {

	ans := 0
	for _, num := range nums {

		temp := 0
		sum := 0
		
		for i := 1; i * i <= num; i++ {
			if num % i == 0 {

				temp++
				sum += i

				if i != num / i {
					temp++
					sum += num / i
				}
			}
		}

		if temp == 4 {
			ans += sum
		}
	}

	return ans
}
