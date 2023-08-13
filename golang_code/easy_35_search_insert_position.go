/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

package golangcode

// @lc code=start
//
//lint:ignore U1000 //
func searchInsert(nums []int, target int) int {
    l := -1
	r := len(nums)
	var m int
	for r - l > 1 {
		m = (l + r) / 2 
		if nums[m] >= target {
			r = m
		} else {
			l = m
		}
	}	
	return r
}
// @lc code=end

