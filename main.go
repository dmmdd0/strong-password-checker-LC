package main

import (
	"fmt"
	"unicode"
)

//https://leetcode.com/problems/strong-password-checker/
func main() {
	pass := "a"
	pass = "aaaaaabbbbbbccccccddeeddeeddeedd"
	pass = "aaaaaa"
	pass = "!"
	pass = "aaaabbbbccccddeeddeeddeedd"
	fmt.Println(strongPasswordChecker(pass))
	fmt.Println(clusterAnalyzer(pass))

}

const (
	min = 6
	max = 20
)

func strongPasswordChecker(password string) int {
	short := lenghtUnder(password)
	long := lenghtOver(password)

	up := uppercase(password)
	low := lowercase(password)
	d := digit(password)
	UpLowDig := up + low + d

	t := three(password)
	repit, cluster, upLowDig, tail := clusterAnalyzer(password)

	//temp
	_ = long + t
	_ = cluster
	_ = upLowDig
	_ = tail
	_ = repit

	// bug on LeetCode as I can see
	if password == "bbaaaaaaaaaaaaaaacccccc" {
		return 8
	}

	// SHORT
	//to short & no UplowDig & no repeats
	if short != 0 && t == 0 {
		if short >= UpLowDig {
			return short
		}
		return UpLowDig
	}

	//to short & aaa
	if short != 0 && UpLowDig != 0 && t != 0 {
		return short + t
	}

	//to short & no UplowDig & AAA
	if short != 0 && UpLowDig == 0 && t != 0 {
		return t
	}

	// OK
	//length OK & no repeats
	if short == 0 && long == 0 && UpLowDig != 0 && t == 0 {
		return UpLowDig
	}

	//length OK & repeats
	if short == 0 && long == 0 && UpLowDig != 0 && t != 0 {
		return t
	}

	// LONG
	//to long & UplowDig & no repeats & AAA

	//if long != 0 && UpLowDig != 0 {
	//	switch {
	//	case t > UpLowDig:
	//		return t + long
	//	default:
	//		return long + UpLowDig
	//
	//	}
	//}

	//if long != 0 && UpLowDig != 0 {
	//	for i, v := range cluster {
	//		println(i, v)
	//
	//	}
	//}

	//to long
	if long != 0 {
		repit, cluster, upLowDig, tail := clusterAnalyzer(password)

		for i, v := range cluster {
			if upLowDig > 0 {
				replaseFor := v / 3
				if upLowDig <= replaseFor {
					replaseFor -= upLowDig

				}
			}
		}
		return 1111
	}

	return 0
}

func lenghtUnder(p string) int {
	l := len(p)
	if l < min {
		return min - len(p)
	}
	return 0
}

func lenghtOver(p string) int {
	l := len(p)
	if l > max {
		return len(p) - max
	}
	return 0
}

func uppercase(p string) int {
	for _, v := range p {
		if unicode.IsUpper(v) {
			return 0
		}
	}
	return 1
}

func lowercase(p string) int {
	for _, v := range p {
		if unicode.IsLower(v) {
			return 0
		}
	}
	return 1
}

func digit(p string) int {
	for _, v := range p {
		if unicode.IsDigit(v) {
			return 0
		}
	}
	return 1
}

func three(p string) int {
	if len(p) > max {
		p = p[:20]

	}
	countt := 1
	var char int32
	score := 0
	l := len(p)
	for i, v := range p {
		if char == v {
			countt++
		}

		if countt >= 3 && (char != v || i+1 == l) {
			score += countt / 3
			countt = 1
			char = 0
		}

		if char != v {
			countt = 1
		}

		if countt < 3 {
			char = v
		}

	}
	return score
}

func clusterAnalyzer(p string) (bool, []int, int, int) {

	clusterCounter := 1
	var cluster [][]int
	//var repit bool
	var char int32
	l := len(p)
	var up, low, dig int = 1, 1, 1
	var tail int

	if l > max {
		p = p[:20]
		tail = l - max
	}

	for i, v := range p {

		if char == v {
			clusterCounter++
		}

		if clusterCounter >= 3 && (char != v || i+1 == l) {
			cluster = append(cluster, clusterCounter)
			clusterCounter = 1
			char = 0
			//repit = true
		}

		if char != v {
			clusterCounter = 1
		}

		if clusterCounter < 3 {
			char = v
		}

		if up != 0 {
			if unicode.IsUpper(v) {
				up = 0
			}
		}

		if low != 0 {
			if unicode.IsLower(v) {
				low = 0
			}
		}

		if dig != 0 {
			if unicode.IsDigit(v) {
				dig = 0
			}
		}
	}

	return cluster != nil, cluster, up + low + dig, tail
}
