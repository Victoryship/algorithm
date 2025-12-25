package dfs

import (
	"sort"
	"strings"
)

/*
17. 电话号码的字母组合
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

	示例 1：
		输入：digits = "23"
		输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

	示例 2：
		输入：digits = "2"
		输出：["a","b","c"]

	示例 3：
		输入：digits = "9"
		输出：["w","x","y","z"]
*/
func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	var (
		phoneMap = map[string]string{
			"2": "abc",
			"3": "def",
			"4": "ghi",
			"5": "jkl",
			"6": "mno",
			"7": "pqrs",
			"8": "tuv",
			"9": "wxyz",
		}
		res       = make([]string, 0)
		backtrack func(index int, choose string)
	)
	backtrack = func(index int, choose string) {
		if index == len(digits) {
			res = append(res, choose)
			return
		}

		letters := string(digits[index])
		for _, v := range phoneMap[letters] {
			backtrack(index+1, choose+string(v))
		}
	}

	backtrack(0, "")
	return res
}

func LetterCombinations1(n int) []string {
	if n == 0 {
		return []string{}
	}

	var (
		res         = make([]string, 0)
		backtrack   func(index int, choose string)
		selecSlices = []string{"0", "1"}
	)
	backtrack = func(index int, choose string) {
		if index == n {
			res = append(res, choose)
			return
		}

		for _, v := range selecSlices {
			backtrack(index+1, choose+string(v))
		}
	}

	backtrack(0, "")
	return res
}

/*
46. 全排列
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

	示例 1：
		输入：nums = [1,2,3]
		输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

	示例 2：
		输入：nums = [0,1]
		输出：[[0,1],[1,0]]

	示例 3：
		输入：nums = [1]
		输出：[[1]]
*/
func Permute(nums []int) [][]int {
	var (
		res, length = [][]int{}, len(nums)
		path        = make([]int, 0, length)
		isUsed      = make([]bool, length)
		backtrack   func()
	)

	backtrack = func() {
		if length == len(path) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for k, v := range nums {
			if isUsed[k] {
				continue
			}

			isUsed[k] = true
			path = append(path, v)
			backtrack()
			isUsed[k] = false
			path = path[:len(path)-1]
		}
	}

	backtrack()
	return res
}

/*
47. 全排列 II
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
	示例 1：
		输入：nums = [1,1,2]
		输出：[[1,1,2],[1,2,1],[2,1,1]]

	示例 2：
		输入：nums = [1,2,3]
		输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

思路： 对于选择树同一层级相同数字的数字结果是一样重复的可以剪枝去重。 所以先对nums进行排序。 元素和前一个元素相同，且前一个元素未被访问（说明是同一层的重复选择）
	visited[i-1] = false：表示前一个相同元素在 “同一层” 未被选择，此时选当前元素会和选前一个元素生成重复排列；
	visited[i-1] = true：表示前一个相同元素在 “上一层” 已被选择，属于合法的不同层选择（比如 nums=[1,1,2]，第一层选第一个 1，第二层选第二个 1，是合法的 [1,1,2]）。
*/

func PermuteUnique(nums []int) [][]int {
	var (
		res, length = [][]int{}, len(nums)
		path        = make([]int, 0, length)
		isUsed      = make([]bool, length)
		backtrack   func()
	)
	// 先进行排序
	sort.Ints(nums)
	backtrack = func() {
		if length == len(path) {
			temp := make([]int, length)
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for k, v := range nums {
			if isUsed[k] {
				continue
			}

			// 条件2：同一层重复元素，跳过（核心剪枝）
			if k > 0 && nums[k] == nums[k-1] && !isUsed[k-1] {
				continue
			}

			isUsed[k] = true
			path = append(path, v)
			backtrack()
			isUsed[k] = false
			path = path[:len(path)-1]
		}
	}

	backtrack()
	return res
}

/*
51. N 皇后
按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
	示例 1：
		输入：n = 4
		输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]

	示例 2：
		输入：n = 1
		输出：[["Q"]]
*/

func SolveNQueens(n int) [][]string {
	var (
		res       = make([][]string, 0, n)
		backtrack func(row int)
		choose    = make([]string, n)
	)

	for i := 0; i < n; i++ {
		choose[i] = strings.Repeat(".", n)
	}

	backtrack = func(row int) {
		if row == n {
			temp := make([]string, n)
			copy(temp, choose)
			res = append(res, temp)
			return
		}

		for i := 0; i < n; i++ {
			if !isVaild(row, i, choose) {
				continue
			}
			str := []rune(choose[row])
			str[i] = 'Q'
			choose[row] = string(str)
			backtrack(row + 1)
			str[i] = '.'
			choose[row] = string(str)
		}
	}

	backtrack(0)
	return res
}

func isVaild(row, col int, choose []string) bool {
	// 检查左斜边是否有“Q”
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if choose[i][j] == 'Q' {
			return false
		}
	}

	// 检查右斜边是否有“Q”
	for i, j := row-1, col+1; i >= 0 && j < len(choose); i, j = i-1, j+1 {
		if choose[i][j] == 'Q' {
			return false
		}
	}

	// 检查同一列是否有“Q”
	for i := row; i >= 0; i-- {
		if choose[i][col] == 'Q' {
			return false
		}
	}

	return true
}
