/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

package golangcode

// @lc code=start
func twoSum(nums []int, target int) []int {
    // m := make(map[int]int) 
	ans := []int{}
	check := false
	for i := 0; i < len(nums)-1; i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				ans = append(ans, i)
				ans = append(ans, j)
				check = true
				break
			}
		}
		if check {
			break
		}
	}
	return ans
}
// @lc code=end

