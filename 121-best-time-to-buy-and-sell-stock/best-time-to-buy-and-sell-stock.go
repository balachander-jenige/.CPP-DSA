func maxProfit(prices []int) int {

    buyprice := prices[0]
    maxprofit :=0

    for _,num := range prices{
        if buyprice > num{
            buyprice = num
        }
        maxprofit = max(maxprofit, num - buyprice)
    }
    return maxprofit
    
}