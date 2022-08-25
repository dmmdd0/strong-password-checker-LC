package main

import (
	"fmt"
	"unicode"
)

//https://leetcode.com/problems/strong-password-checker/
func main() {
	fmt.Println(strongPasswordChecker("bbaa1aa2aa3aa4aa5ccA123"))
	//	"bbaaaaaaaaaaaaaaacccccc"
	//	"bbaaaaaaaaaaaaaaaccc ccc"
	//	"bb aaa aaa aaa aaa aaa ccc / ccc"

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

	t := three(password)

	//temp
	_ = long + t

	UpLowDig := up + low + d

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

	//length OK & no repeats
	if short == 0 && long == 0 && UpLowDig != 0 && t == 0 {
		return UpLowDig
	}

	//length OK & repeats
	if short == 0 && long == 0 && UpLowDig != 0 && t != 0 {
		return t
	}

	//to long & no repeats
	//to long & aaa
	if long != 0 && UpLowDig != 0 {
		return long + UpLowDig
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
