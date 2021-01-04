package dp

import (
	. "dp/tools"
	"fmt"
	"math"
)

func MaxProfitK1(k int, prices []int) int {
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

func StockDp(prices []int) (s int) {
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

func StockGreed(prices []int) (ans int) {
	for i := 1; i < len(prices); i++ {
		ans += Max(0, prices[i]-prices[i-1])
	}
	return
}

func MaxProfit1(prices []int) int {
	mi, ma := INT_MAX, 0
	for _, price := range prices {
		ma = Max(price-mi, ma)
		mi = Min(price, mi)
	}
	return ma
}

func MaxProfit2(prices []int) int {
	mi, ma := int(1e9), 0
	for _, v := range prices {
		if v < mi {
			mi = v
		} else if v-mi > ma {
			ma = v - mi
		}
	}
	return ma
}

func Stock(prices []int) {
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

func ImprovedStock(prices []int) {
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