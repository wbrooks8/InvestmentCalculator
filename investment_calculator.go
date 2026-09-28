package main

import (
	"bufio"
	"errors"
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

	investmentAmount, err := handleInput("investmentAmount")

	if err != nil {
		fmt.Println("Input failed. Please try again.")
		return
	}

	years, err = handleInput("years")

	if err != nil {
		fmt.Println("Input failed. Please try again.")
		return
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

func handleInput(version string) (float64, error) {

	var value float64
	var err error

	scanner := bufio.NewScanner(os.Stdin)

	for {
		switch version {
		case "investmentAmount":
			fmt.Print("Please give your investment amount: ")
		case "years":
			fmt.Print("Please give the years of the investment: ")
		}

		if !scanner.Scan() {
			err := errors.New("Input failed. Please try again.")
			return 0.0, err
		}

		value, err = strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)

		if err != nil || value <= 0 {
			if version == "investmentAmount" {
				fmt.Println("Please enter a number greater than 0.")
				continue
			} else if version == "years" {
				fmt.Println("Please enter a number greater than 0 and less than or equal to 100")
				continue
			}
		}

		break
	}

	if err := scanner.Err(); err != nil {
		return 0.0, err
	}

	return value, nil
}
