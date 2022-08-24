package main

import (
	"fmt"
	"unicode"
)

//https://leetcode.com/problems/strong-password-checker/
func main() {

	fmt.Println(strongPasswordChecker("ABABABABABABABABABAB1"))

}

const (
	min = 6
	max = 20
)

func strongPasswordChecker(password string) int {
	l := lenght(password)
	if l != 0 {
		return l
	}

	t := three(password)
	if t > 0 {
		return t + digit(password)
	}

	return l +
		uppercase(password) +
		lowercase(password) +
		digit(password)
}

func lenght(p string) int {
	l := len(p)
	if l < min {
		return min - l
	}
	if l > max {
		return l - max
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
