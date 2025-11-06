package main

func smallestNumberWithAllSetBits(n int) int {
	if n == 0 {
		return 0
	}

	result := 0
	bitPosition := 0

	for n > 0 {
		if n&1 == 1 {
			result += 1 << bitPosition
		}
		bitPosition++
		n >>= 1
	}

	return result
}

// func main() {
// 	// Example usage
// 	n := 10
// 	println(smallestNumberWithAllSetBits(n)) // Output: 15
// }
