package random

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Function to calculate maximum profit by buying and selling a stock from an array of stock prices for a day.
// The indices are the time (in minutes) past trade opening time, which was 9:30 am local time.
// The values are the price (in US dollars) of one share of certain stock at that time.So if the stock cost $500 at 10:30am,
// that means stockPrices[60] = 500.
// Constraint - You need to buy before you can sell. Also, you can't buy and sell in the same time step—at least 1 minute has to pass.
// Runtime Complexity - TIme: O(n), Space: O(1)
func getMaxProfit(prices []int) int {
	var n, maxProfit, tempProfit, minPrice int
	n = len(prices)

	if n < 2 {
		panic("There has to at least 2 price to have a profit")
	}

	// we'll greedily update minPrice and maxProfit, so we initialize
	// them to the first price and the first possible profit
	minPrice = prices[0]
	maxProfit = prices[1] - prices[0]

	// start at the second (index 1) time we can't sell at the first time, since we must buy first,
	// and we can't buy and sell at the same time!
	// if we started at index 0, we'd try to buy *and* sell at time 0. this would give a profit of 0, which is a problem if our
	// maxProfit is supposed to be *negative*--we'd return 0.
	for i := 1; i < n; i++ {

		// see what our profit would be if we bought at the min price and sold at the current price
		tempProfit = prices[i] - minPrice

		// update maxProfit if we can do better
		maxProfit = max(maxProfit, tempProfit)

		// update minPrice so it's always the lowest price we've seen so far
		minPrice = min(minPrice, prices[i])

	}
	return maxProfit

}
