package main

import (
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	ans := make([]int, n)
	stillRun := 0
	for i := 0; i < len(logs); i++ {
		cLogs := strings.Split(logs[i], ":")
		f := cLogs[0]
		startEnd := cLogs[1]
		time := cLogs[2]
		
		if i == 0 {
			stillRun = 0
		}

		fNum, _ := strconv.Atoi(f)
		nTime, _ := strconv.Atoi(time)

		switch startEnd {
		case "start":
			if stillRun != i {
				ans[stillRun] = nTime - ans[fNum]
			}
			ans[fNum] = nTime
		case "end":
			ans[fNum] = nTime - ans[fNum] + 1
		}
	}
	return ans
}
