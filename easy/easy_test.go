package main

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	twoSum := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(twoSum)
}

func TestIsPalindrome(t *testing.T) {
	isPalindrome := isPalindrome(1221)
	fmt.Println(isPalindrome)
}

func TestRomanToInt(t *testing.T) {
	romanToInt := romanToInt("MCMXCIV")
	fmt.Println(romanToInt)
}

func TestLongestCommonPrefix(t *testing.T) {
	longestCommonPrefix := longestCommonPrefix([]string{"flower", "flow", "flight"})
	fmt.Println(longestCommonPrefix)
}

func TestIsValid(t *testing.T) {
	isValid := isValid("({})[{}]({([])})")
	fmt.Println(isValid)
}
