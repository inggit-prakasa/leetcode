package main

// Bruteforce
func twoEditWords(queries []string, dictionary []string) []string {
	res := make([]string, 0)

	for _, word := range queries {
		for _, dict := range dictionary {
			temp := 0
			for i := 0; i < len(dict); i++ {
				if word[i] != dict[i] {
					temp++
				}
			}

			if temp <= 2 {
				res = append(res, word)
				break
			}
		}
	}

	return res
}
