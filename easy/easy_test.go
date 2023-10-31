package main

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	twoSum := TwoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(twoSum)
}

func TestIsPalindrome(t *testing.T) {
	isPalindrome := IsPalindrome(1221)
	fmt.Println(isPalindrome)
}

func TestRomanToInt(t *testing.T) {
	romanToInt := RomanToInt("MCMXCIV")
	fmt.Println(romanToInt)
}
