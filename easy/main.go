package main

import "fmt"

// 1. TwoSums, Using Hashtable solution
func twoSum(nums []int, target int) []int {
	//Making hash table
	/*
		Hashtable :
		Index (the value of current nums will be the index)
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
		//if the index reqNums is not found in hash maps we put the value of current nums as index in hash maps with the current index of nums as value
		hashMaps[v] = i
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

// 14. Longest Common Prefix
func longestCommonPrefix(strs []string) string {
	//if there is only 1 data in the array return the data because the prefix is that data
	if len(strs) == 1 {
		return strs[0]
	}
	//looping the first array as base string for the prefix
	for i, v := range strs[0] {
		//looping the array except the first one (index 0)
		for j := 1; j <= len(strs)-1; j++ {
			//if base string  index is larger or equal to the length of string of current string return the base string slice up to i from beginning
			//if the current string slice i is not equal to base string value byte return the base string slice up to i from beginning
			if i >= len(strs[j]) || strs[j][i] != byte(v) {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

// 20. Valid Parentheses
func isValid(s string) bool {
	//if s has odd character return false
	if len(s)%2 != 0 {
		return false
	}
	//if the 1st character of s is close bracket return false
	if s[0] == ')' || s[0] == '}' || s[0] == ']' {
		return false
	}
	var openBracket []byte
	//closeToOpening to get open bracket using close bracket
	closeToOpening := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	//loop s
	for _, v := range s {
		//if byte(v) is in closeToOpening its a open bracket otherwise its close bracket
		if isOpen, isClosed := closeToOpening[byte(v)]; !isClosed {
			//if its a open bracket append it to openBracket
			openBracket = append(openBracket, byte(v))
		} else {
			//if its not a open bracket
			//we check if open bracket is empty if it is empty return false
			//we check again if the last array of openBracket is not equal to isOpen return false
			if len(openBracket) == 0 || openBracket[len(openBracket)-1] != isOpen {
				return false
			}
			//removing last array of openBracket because we found the matching pair for the close bracket
			openBracket = openBracket[:len(openBracket)-1]
		}
	}
	//finally, we check if there is still left over openBracket
	return len(openBracket) == 0
}

// 21. Merge Two Sorted Lists, Rewiring Next Solution
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	//if list 1 is nil then list 2 will guarantee be larger than list 1
	if list1 == nil {
		return list2
	}
	//if list 2 is nil then list 1 will guarantee be larger than list 2
	if list2 == nil {
		return list1
	}
	//if list 1 value below list 2 value, we keep the current list 1 value then change list 1 next value calling the function again with list 1 next and list 2 as parameter
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		//if list 2 value below list 1 value, we keep the current list 2 value then change list 2 next value calling the function again with list 2 next and list 1 as parameter
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}

	/*
		Extra Explanation
		Consider we have two linked list like this
		List 1 : 1->2->3->5
		List 2 : 1->3->4

		First, we check list 1 value which is 1 with list 2 value which is also 1, Whatever value is below the other value we gonna keep it, in this case is any of them so we keep list 1 value

		Second, Now since we keeping list 1 value now we go to the next value of list 1

		List 1 : 1(keep)  2(current)->3->5
		List 2 : 1(current)->3->4

		Third, we now checking again both current value of the list, in this case List 2 value of 1 is win, so now we keep this value and go to the next value of list 2

		List 1 : 1  2(current)->3->5
						 ↓
		List 2 : 1(keep) 3(current)->4

		We know see how this is work, by comparing each current value and move the keep to the lowest value we compare, okay lets move to next step

		Fourth,  we now checking again both current value of the list, in this case List 1 value of 2 is win, so now we keep this value and go to the next value of list 1

		List 1 : 1  2(keep) 3(current)->5
						 ↓  ↑
		List 2 : 1 -> 3(current)->4

		Fifth,  we now checking again both current value of the list, in this case its a tie so we keep List 1 value of 3, so now we keep this value and go to the next value of list 1

		List 1 : 1  2->3(keep) 5(current)
						 ↓  ↑
		List 2 : 1 -> 3(current)->4

		Sixth,  we now checking again both current value of the list, in this case List 2 value of 3 is win, so now we keep this value and go to the next value of list 2

		List 1 : 1  2-> 3 5(current)
						 ↓  ↑  ↓
		List 2 : 1 -> 3(keep) 4(current)

		Seventh, we now checking again both current value of the list, in this case List 2 value of 4 is win, so now we keep this value and go to the next value of list 2

		List 1 : 1  2-> 3 5(current)
						 ↓  ↑  ↓
		List 2 : 1 -> 3->4

		Eight, List 2 already reach the end of the linked list, since we already "rewiring" each value of the linked list and keeping only the lowest we can guarantee that List 1 and beyond will always be higher than List 2, so if any of the List reach null we can just return the other list

		Hopefully, Me (in the future) and You understand how it works, its hard to explain in comment but theres always youtube video

	*/
}

func main() {
	// twoSums := twoSum([]int{2, 7, 11, 15}, 9)
	// IsPalindrome := isPalindrome(2222)
	// RomanToInt := romanToInt("MDCCCLXXXIV")
	// LongestCommonPrefix := longestCommonPrefix([]string{"abc", "ab", "abc"})
	// isValid := isValid("()")
	list1 := &ListNode{Val: 2, Next: nil}
	list2 := &ListNode{Val: 1, Next: nil}
	mergeTwoLists := mergeTwoLists(list1, list2)
	fmt.Println(mergeTwoLists)
}
