/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

package golangcode

// @lc code=start
//
//lint:ignore U1000 //
func plusOne(digits []int) []int {
	lastSize := len(digits) - 1
    if digits[lastSize] < 9 {
		digits[lastSize] = digits[lastSize] + 1
	} else {
		ans := []int{1}
		bcnt := make([]bool, len(digits))
		// 与えられた配列の要素の初期値が9かどうか
		for i, v := range digits {
			if v == 9 {
				bcnt[i] = true
			}
		}
		for i := lastSize; i >= 0; i-- {
			if bcnt[i] {
				digits[i] = 0
				if i == 0 {
					ans = append(ans, digits...)
					return ans
				}
			} else {
				digits[i] = digits[i] + 1
				return digits
			}
		}
	}
	return digits
}
// @lc code=end

