package main

import (
	"fmt"
	"unicode"
)

//https://leetcode.com/problems/strong-password-checker/
func main() {

	fmt.Println(strongPasswordChecker("aaabb1"))

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

	//to short
	if short != 0 && UpLowDig != 0 && t == 0 {
		if short >= UpLowDig {
			return short
		}
		return UpLowDig
	}

	//to short & aaa
	if short != 0 && UpLowDig != 0 && t != 0 {
		return short + t
	}
	return 11111111
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
