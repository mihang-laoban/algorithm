package dp

import (
	. "dp/tools"
	"fmt"
	"math"
	"testing"
)

/*示例1：
输入：[3, 3, 5, 0, 0, 3, 1, 4]
输出：6
解释：在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3 - 0 = 3 。随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4 - 1 = 3 。
示例2：
输入：[1, 2, 3, 4, 5]
输出：4
解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4。需要注意的是，你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
示例3：
输入：[7, 6, 4, 3, 1]
输出：0
解释：在这个情况下, 没有交易完成, 所以最大利润为 0。*/
// Q17
func TestStock(t *testing.T) {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	length := len(prices)
	dp := make([][2][3]int, length)
	// 假设第一天没有买入
	dp[0][0][0] = 0
	dp[0][0][1] = 0
	dp[0][0][2] = 0
	// 第一天不可能已卖出
	dp[0][1][0] = -prices[0]
	dp[0][1][1] = -prices[0]
	dp[0][1][2] = -prices[0]
	// 处理后续日期
	for i := 1; i < length; i++ {
		dp[i][0][0] = 0
		dp[i][0][1] = Max(dp[i-1][1][0]+prices[i], dp[i-1][0][1])
		dp[i][0][2] = Max(dp[i-1][1][1]+prices[i], dp[i-1][0][2])
		dp[i][1][0] = Max(dp[i-1][0][0]-prices[i], dp[i-1][1][0])
		dp[i][1][1] = Max(dp[i-1][0][1]-prices[i], dp[i-1][1][1])
		dp[i][1][2] = 0
	}
	res := Max(dp[length-1][0][1], dp[length-1][0][2])
	fmt.Println(res)
}

func TestStock2(t *testing.T) {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	size := len(prices)

	dp := [2][2][3]int{}
	dp[0][0][0] = 0
	dp[0][0][1] = 0
	dp[0][0][2] = 0

	dp[0][1][0] = -prices[0]
	dp[0][1][1] = -prices[0]
	dp[0][1][2] = -prices[0]

	for i := 1; i < size; i++ {
		dp[i%2][0][0] = 0
		dp[i%2][0][1] = Max(dp[(i-1)%2][1][0]+prices[i], dp[(i-1)%2][0][1])
		dp[i%2][0][2] = Max(dp[(i-1)%2][1][1]+prices[i], dp[(i-1)%2][0][2])
		dp[i%2][1][0] = Max(dp[(i-1)%2][0][0]-prices[i], dp[(i-1)%2][1][0])
		dp[i%2][1][1] = Max(dp[(i-1)%2][0][1]-prices[i], dp[(i-1)%2][1][1])
		dp[i%2][1][2] = 0
	}
	res := Max(dp[(size-1)%2][0][1], dp[(size-1)%2][0][2])
	fmt.Println(res)
}

/*给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。
注意：你不能在买入股票前卖出股票。

示例 1:
输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。

示例 2:
输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
// Q18
func TestMaxProfit1(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit2(prices))
}

func maxProfit(prices []int) int {
	mi := INT_MAX
	//mi := int(1e9)
	ma := 0
	for _, price := range prices {
		ma = Max(price-mi, ma)
		mi = Min(price, mi)
	}
	return ma
}

func maxProfit2(prices []int) int {
	mi := int(1e9)
	ma := 0

	for _, v := range prices {
		if v < mi {
			mi = v
		} else if v-mi > ma {
			ma = v - mi
		}
	}

	return ma
}

/*给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:
输入: [7,1,5,3,6,4]
输出: 7
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。

示例 2:
输入: [1,2,3,4,5]
输出: 4
解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
     因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。

示例 3:
输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
// Q19
func TestMaxProfit2(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(stockDp(prices))
}

func stockDp(prices []int) (s int) {
	s = 0
	e := -prices[0]
	for i := 1; i < len(prices); i++ {
		s = Max(s, e+prices[i])
		e = Max(e, s-prices[i])
	}
	return
	/*	dp := make([][2]int, size)
		dp[0][0] = 0
		dp[0][1] = -prices[0]
		for i := 1; i < size; i++ {
			dp[i][0] = Max(dp[i-1][1]+prices[i], dp[i-1][0])
			dp[i][1] = Max(dp[i-1][0]-prices[i], dp[i-1][1])
		}
		return dp[size-1][0]
	*/
}

func stockGreed(prices []int) (ans int) {
	for i := 1; i < len(prices); i++ {
		ans += Max(0, prices[i]-prices[i-1])
	}
	return
}

/*给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。



示例 1：

输入：k = 2, prices = [2,4,1]
输出：2
解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
示例 2：

输入：k = 2, prices = [3,2,6,5,0,3]
输出：7
解释：在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3 。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
// Q20
func TestMaxProfitK(t *testing.T) {
	fmt.Println(maxProfitK1(2, []int{3, 2, 6, 5, 0, 3}))
	//fmt.Println(maxProfitK1(2, []int{3, 3, 5, 0, 0, 3, 1, 4}))
}

func maxProfitK1(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	k = Mini(k, n/2)

	buy := make([]int, k+1)
	sell := make([]int, k+1)

	buy[0] = -prices[0]
	for i := 1; i <= k; i++ {
		buy[i] = math.MinInt64 / 2
		sell[i] = math.MinInt64 / 2
	}

	for i := 1; i < n; i++ {
		buy[0] = Maxi(buy[0], sell[0]-prices[i])
		for j := 1; j <= k; j++ {
			buy[j] = Maxi(buy[j], sell[j]-prices[i])
			sell[j] = Maxi(sell[j], buy[j-1]+prices[i])
		}
	}
	return Maxi(sell...)
}
