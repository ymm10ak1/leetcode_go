/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

package golangcode

import "strconv"

// @lc code=start
//
//lint:ignore U1000 //
func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	rs := []rune(s)
	size := len(s)-1
	check := true
	for i := 0; i < len(s); i++ {
		if rs[i] != rs[size-i] {
			check = false
		}
	}
	return check
}
// @lc code=end

