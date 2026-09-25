package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const InflationRate = 2.5
const ExpectedReturnRate = 5.5

func main() {
	var investmentAmount float64
	var years float64
	scanner := bufio.NewScanner(os.Stdin)

	// "&" creates a pointer to the variable so it can be reassigned.
	for {
		fmt.Print("Please give your investment amount: ")

		scanner.Scan()
		value, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)

		if err != nil || value <= 0 {
			fmt.Println("Please enter a number greater than 0.")
			continue
		}

		investmentAmount = value
		break
	}

	for {
		fmt.Println("Please give the number of years")

		scanner.Scan()
		value, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)

		if err != nil || value <= 0 || value > 100 {
			fmt.Println("Please enter a number greater than 0 and less than 100")
			continue
		}

		years = value
		break
	}

	futureValue := calculateFutureValue(investmentAmount, years)
	// Calculating with inflation
	futureRealValue := calculateRealFutureValue(futureValue, years)

	// fmt.Sprintf allows the creation of formatted string that can be stored as var
	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (Adjusted for Inflation): %.2f\n", futureRealValue)

	fmt.Print(formattedFV)
	fmt.Print(formattedRFV)

}

func calculateFutureValue(amount, years float64) float64 {
	return amount * math.Pow(1+ExpectedReturnRate/100, years)
}

func calculateRealFutureValue(amount, years float64) float64 {
	return amount / math.Pow(1+InflationRate/100, years)
}
