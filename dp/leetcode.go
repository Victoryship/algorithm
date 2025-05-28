package dp

/*
*
121. 买卖股票的最佳时机
	给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
	你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
	返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

	示例 1：
	输入：[7,1,5,3,6,4]
	输出：5
	解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
		 注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。

	示例 2：
	输入：prices = [7,6,4,3,1]
	输出：0
	解释：在这种情况下, 没有交易完成, 所以最大利润为 0。
*/

// 贪心算法思路：一个变量记录股票的最低价格，遍历数组时计算当前价格与最低价格的差值，更新最大利润。
func MaxProfit(prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}

	minPrice, res := prices[0], 0
	for _, v := range prices {
		if v <= minPrice {
			minPrice = v
			continue
		}

		if res < v-minPrice {
			res = v - minPrice
		}
	}
	return res
}

/*
动态规划思路：

	1、 定义动态状态： dp[i][0]表示第i天不持有股票的最大利润，dp[i][1]表示第i天持有股票的最大利润
	2、 状态转移方程：
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]) // 第i天不持有股票最大收益为前一天就不持有股票和前一天持有股票并在今天卖出的最大值
		dp[i][1] = max(dp[i-1][1], -prices[i]) // 第i天持有股票的最大收益为前一天就持有股票和今天买入股票的最大值
	3、 状态初始条件：
		dp[0][0] = 0 // 第一天不持有股票的最大收益为0
		dp[0][1] = -prices[0] // 第一天持有股票的最大收益为第一天的价格。
	4、 最终结果为最后一天不吃有股票的最大收益，即dp[i][0]
	5、 动态规划状态数组优化。 dp[i][0] => sell dp[i][1] => hold
*/
func MaxProfitDP(prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}
	hold, sell := -prices[0], 0
	for i := 1; i < length; i++ {
		if sell < hold+prices[i] {
			sell = hold + prices[i]
		}

		if hold < -prices[i] {
			hold = -prices[i]
		}
	}

	return sell
}

/*
*
121. 买卖股票的最佳时机
	给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
	在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
	返回 你能获得的 最大 利润 。

	示例 1：

	输入：prices = [7,1,5,3,6,4]
	输出：7
	解释：在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4。
	随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6 - 3 = 3。
	最大总利润为 4 + 3 = 7 。
	示例 2：

	输入：prices = [1,2,3,4,5]
	输出：4
	解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4。
	最大总利润为 4 。
	示例 3：

	输入：prices = [7,6,4,3,1]
	输出：0
	解释：在这种情况下, 交易无法获得正利润，所以不参与交易可以获得最大利润，最大利润为 0。
*/

// 贪心算法思路：// 在每一天，如果今天的价格大于昨天的价格，则可以卖出股票，计算利润并累加。
func MaxProfit2(prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}

	res := 0
	for i := 1; i < length; i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}

	return res
}

/*
动态规划思路：

	1、 定义动态状态： dp[i][0]表示第i天不持有股票的最大利润，dp[i][1]表示第i天持有股票的最大利润
	2、 状态转移方程：
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]) // 第i天不持有股票最大收益为前一天就不持有股票和前一天持有股票并在今天卖出的最大值
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]) // 第i天持有股票的最大收益为前一天就持有股票和今天买入股票的最大值
	3、 状态初始条件：
		dp[0][0] = 0 // 第一天不持有股票的最大收益为0
		dp[0][1] = -prices[0] // 第一天持有股票的最大收益为第一天的价格。
	4、 最终结果为最后一天不吃有股票的最大收益，即dp[i][0]
	5、 动态规划状态数组优化。 dp[i][0] => sell dp[i][1] => hold
*/
func MaxProfitDP2(prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}
	hold, sell := -prices[0], 0
	for i := 1; i < length; i++ {
		if sell < hold+prices[i] {
			sell = hold + prices[i]
		}

		if hold < sell-prices[i] {
			hold = sell - prices[i]
		}
	}

	return sell
}
