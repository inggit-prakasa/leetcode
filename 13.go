package main

func romanToInt(s string) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case "I":
			if i+1 < len(s) {
				if string(s[i+1]) == "V" {
					ans += 4
					i++
					continue
				}

				if string(s[i+1]) == "X" {
					ans += 9
					i++
					continue
				}
			}
			ans += 1
		case "V":
			ans += 5
		case "X":
			if i+1 < len(s) {
				if string(s[i+1]) == "L" {
					ans += 40
					i++
					continue
				}

				if i+1 < len(s) && string(s[i+1]) == "C" {
					ans += 90
					i++
					continue
				}
			}
			ans += 10
		case "L":
			ans += 50
		case "C":
			if i+1 < len(s) {
				if string(s[i+1]) == "D" {
					ans += 400
					i++
					continue
				}

				if i+1 < len(s) && string(s[i+1]) == "M" {
					ans += 900
					i++
					continue
				}
			}
			ans += 100
		case "D":
			ans += 500
		case "M":
			ans += 1000

		}
	}

	return ans
}
