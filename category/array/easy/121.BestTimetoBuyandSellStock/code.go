package leetcode

/*

使用動態規劃 DP
*/

func maxProfit_new(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	min, maxProfit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > maxProfit {
			maxProfit = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return maxProfit
}

//自己想的 沒料會用到 2 n次方
func maxProfit(prices []int) int {

	var maxValue = 0
	var maxProfix = 0
	for i, value := range prices {

		if value > maxValue && i != 0 {
			continue
		}

		for j := i; j < len(prices); j++ {
			if prices[j] > maxValue {
				maxValue = prices[j]
			}

			if prices[j] > value && (prices[j]-value) > maxProfix {
				maxProfix = prices[j] - value
			}
		}

	}
	return maxProfix
}
