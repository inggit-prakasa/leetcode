package main

func countBinarySubstrings(s string) int {
	zero := 0
	one := 0
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			zero++
		} else {
			one++
		}

		if zero == one && zero != 0 && one != 0 {
			ans = ans + ( (one + zero) / 2 )

			if s[i] == '0' {
				one = 0
			} else {
				zero = 0
			}
		}
	}

	return ans
}
