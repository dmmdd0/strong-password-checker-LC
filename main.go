package main

import (
	"fmt"
	"unicode"
)

//https://leetcode.com/problems/strong-password-checker/
func main() {

	fmt.Println(strongPasswordChecker("aabb1"))

}

const (
	min = 6
	max = 20
)

func strongPasswordChecker(password string) int {
	sh := lenghtUnder(password)
	lon := lenghtOver(password)

	up := uppercase(password)
	low := lowercase(password)
	d := digit(password)

	t := three(password)

	//temp
	_ = lon + t

	UpLowDig := up + low + d

	if sh != 0 && UpLowDig != 0 {
		switch sh >= UpLowDig {
		case true:
			return sh
		default:
			return UpLowDig
		}
	}
	return 11111111

}

func lenghtUnder(p string) int {
	return min - len(p)
}

func lenghtOver(p string) int {
	return len(p) - max
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
