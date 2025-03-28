package leetcode

/*
LongestPalindrome 5. 最长回文子串
给你一个字符串s，找到s中最长的回文子串。

示例 1：

	输入：s = "babad"
	输出："bab"
	解释："aba" 同样是符合题意的答案。

示例 2：

	输入：s = "cbbd"
	输出："bb"
*/
func LongestPalindrome(s string) string {
	var (
		length = len(s)
		res    string
	)

	if length <= 1 {
		return s
	}

	for i := 0; i < length; i++ {
		str1 := getPalindrome(s, i, i, length)   // 奇数回文子串
		str2 := getPalindrome(s, i, i+1, length) // 偶数回文子串

		if len(str1) > len(str2) && len(str1) > len(res) {
			res = str1
		}

		if len(str2) > len(str1) && len(str2) > len(res) {
			res = str2
		}
	}

	return res
}

// getPalindrome 获取回文字符串
func getPalindrome(s string, l, r, length int) string {
	for l >= 0 && r < length && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}

/*
MaxArea 11. 盛最多水的容器
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
返回容器可以储存的最大水量。(计算数组在二维坐标上组成的最大矩形的面积)

示例 1：

	输入：[1,8,6,2,5,4,8,3,7]
	输出：49
	解释：垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49

示例 2：

	输入：height = [1,1]
	输出：1
*/
func MaxArea(height []int) int {
	length := len(height)
	if length == 0 {
		return 0
	}

	left, right, res := 0, length-1, height[0]
	for left < right {
		min := height[left]
		if height[right] < min {
			min = height[right]
		}

		temp := min * (right - left)
		if temp > res {
			res = temp
		}

		if height[left] < height[right] {
			left++
			continue
		}

		right--
	}

	return res
}

/*
RemoveDuplicates 26. 删除有序数组中的重复项
给你一个 非严格递增排列 的数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次 ，返回删除后数组的新长度。元素的相对顺序应该保持一致。
然后返回 nums 中唯一元素的个数。

示例 1：

	输入：nums = [1,1,2]
	输出：2, nums = [1,2,_]
	解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。

示例 2：

	输入：nums = [0,0,1,1,1,2,2,3,3,4]
	输出：5, nums = [0,1,2,3,4]
	解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。
*/
func RemoveDuplicates(nums []int) int {
	fast, slow, length := 0, 0, len(nums)
	if length == 0 {
		return 0
	}

	for fast < length {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}

	return slow + 1
}

/*
RemoveElement 27. 移除元素
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素。元素的顺序可能发生改变。然后返回 nums 中与 val 不同的元素的数量。

示例 1：

	输入：nums = [3,2,2,3], val = 3
	输出：2, nums = [2,2,_,_]
	解释：你的函数函数应该返回 k = 2, 并且 nums 中的前两个元素均为 2。
	你在返回的 k 个元素之外留下了什么并不重要（因此它们并不计入评测）。

示例 2：

	输入：nums = [0,1,2,2,3,0,4,2], val = 2
	输出：5, nums = [0,1,4,0,3,_,_,_]
	解释：你的函数应该返回 k = 5，并且 nums 中的前五个元素为 0,0,1,3,4。
	注意这五个元素可以任意顺序返回。
	你在返回的 k 个元素之外留下了什么并不重要（因此它们并不计入评测）。
*/
func RemoveElement(nums []int, val int) int {
	fast, slow, length := 0, 0, len(nums)
	for fast < length {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

/*
Trap 42. 接雨水
给定n个非负整数表示每个宽度为1的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例 1：

	输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
	输出：6
	解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

示例 2：

	输入：height = [4,2,0,3,2,5]
	输出：9
*/
func Trap(height []int) int {
	res, length := 0, len(height)
	if length < 3 {
		return res
	}

	maxLeft, maxRight := height[0], height[length-1]
	// 能接住雨水的数组必须要有低谷，即对于某个位置。 在左右两边找到最大的柱子。 该位置能接的雨水为其中更小的一个值
	for i := 1; i < length; i++ {

	}

	return res
}
