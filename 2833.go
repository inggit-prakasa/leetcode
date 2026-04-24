package main

func furthestDistanceFromOrigin(moves string) int {
	res := 0
	L := 0
	R := 0
	M := 0
	for _, move := range moves {
		switch move {
		case 'L':
			L++
		case 'R':
			R++
		default:
			M++
		}
	}

	res = abs(L-R) + M

	return res
}
