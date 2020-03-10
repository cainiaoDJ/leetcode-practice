/**
https://leetcode-cn.com/problems/coin-change/
零钱兑换问题
*/
package main

import (
	"fmt"
	"sort"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < amount+1; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	sort.Ints(coins)

	for i := 0; i < amount+1; i++ {
		if coins[0] > i {
			continue
		}
		for _, coin := range coins {
			if i < coin {
				continue
			}
			if dp[i] > 1+dp[i-coin] {
				dp[i] = 1 + dp[i-coin]
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}

func main() {
	coins := []int{2, 3, 7, 5}
	fmt.Println(coinChange(coins, 60))
}
