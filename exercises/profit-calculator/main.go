package main

import (
	"fmt"
	"math"
)

func main() {
	grossRevenue := printAndScanFloat("Gross revenue: ")
	expenses := printAndScanFloat("All expenses: ")
	taxRate := printAndScanFloat("Tax rate: ")

	ebt, profit, ratio := calculateProfit(grossRevenue, expenses, taxRate)

	fmt.Printf("EBT: %.2f\n", math.Round(ebt))
	fmt.Printf("Profit: %.2f\n", math.Round(profit))
	fmt.Printf("Profit ratio: %.2f\n", ratio)
}

func printAndScanFloat(strToPrint string) float64 {
	var localFloat float64
	fmt.Print(strToPrint)
	fmt.Scan(&localFloat)

	return localFloat
}

func calculateProfit(grossRevenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := grossRevenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}
