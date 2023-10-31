package main

import "fmt"

// 1. TwoSums, Using Hashtable solution
func twoSum(nums []int, target int) []int {
	//Making hash table
	/*
		Hashtable :
		Index (the required number will be the index)
		Value (the index of current nums will be the value)
	*/
	hashMaps := make(map[int]int)

	//example nums = [2,7,4], target = 9
	for i, v := range nums {
		//we just need to look for the required numbers in hashtable, e.g: current value is 2 then we just need to look for 9-2 = 7, the number 7 in hash table
		reqNums := target - v
		//if the index reqNums is exist in hash maps we return the value of that hash maps (which contain the index nums) and the current index of nums
		if val, exists := hashMaps[reqNums]; exists {
			return []int{val, i}
		}
		//if the index reqNums is not found in hash maps we put the index reqNums in hash maps with the current index of nums as value
		hashMaps[reqNums] = i
	}
	return []int{}
}

// 9. Palindrome Number, Getting each digit of x from last and compare it
func isPalindrome(x int) bool {
	//negative number is guarantee not gonna be palindrome
	if x < 0 {
		return false
	}
	//everything below 10 and above -1 is always palindrome
	if x < 10 {
		return true
	}
	//number cant start from 0, so if it ends with 0 its not palindrome, e.g: 0121 (would just be 121) thus 1210 would never be palindrome because 0121 doesnt exist
	if x%10 == 0 {
		return false
	}

	digits := 0
	i := x
	for i > 0 {
		digits = digits*10 + i%10
		i /= 10
		//Palindrom would always just half of the number digit (skipping the middle value if it odd length number) equal to the other half example 1221 is just 12 (read 2 digit from last) == 12 (read from 1st digit to half), or 121 1 (reading from last digit) == 1 (reading from the first digit) 2 is skipped since its in the middle and its odd length number
		if digits == x || digits == x/10 {
			return true
		}
	}
	return false
}

// 13. Roman to Integer
func romanToInt(s string) int {
	//Roman Numeral Table from 1 to 1000 in map
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	// setting the total to the last roman character of s
	total := roman[s[len(s)-1]]
	//Looping s to get each single character from last - 1 of the string position
	for i := len(s) - 2; i >= 0; i-- {
		//if the currect roman is below the right roman it would substract the current roman e.g (IV, IX), substraction always happen if the left roman is less than the right one
		if roman[s[i]] < roman[s[i+1]] {
			total -= roman[s[i]]
		} else { //if the current roman is not below the right roman it would add the current roman e.g (III, VII), adding always happen if the left roman is higher or equal than the right one
			total += roman[s[i]]
		}
	}
	return total
}

func main() {
	// twoSums := twoSum([]int{3, 2, 4}, 6)
	// isPalindrome := isPalindrome(2222)
	romanToInt := romanToInt("MDCCCLXXXIV")
	fmt.Println(romanToInt)
}