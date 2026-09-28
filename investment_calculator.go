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

var currentValue float64 = 0

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	handleRequest(scanner)

}

func handleRequest(scanner *bufio.Scanner) {
	for {
		fmt.Println("Please tell me what you would like to do...")
		fmt.Println("1. Simulate a single investment")
		fmt.Println("2. Simulate weekly contributions")
		fmt.Println("3. Simulate monthly contributions")
		fmt.Println("4. Simulate yearly contributions")
		fmt.Println("5. Exit the application")

		if !scanner.Scan() {
			panic("Input failed. Please try again.")
		}

		choice, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			fmt.Println("Invalid input. Please enter a number from the menu.")
			continue
		}

		switch choice {
		case 1:
			handleNewSimulation(1, scanner)
		case 2:
			handleNewSimulation(2, scanner)
		case 3:
			handleNewSimulation(3, scanner)
		case 4:
			handleNewSimulation(4, scanner)
		case 5:
			os.Exit(0)
		}
	}

}

func handleNewSimulation(simulationType int, scanner *bufio.Scanner) {
	var investmentAmount float64
	var years float64
	var overTimeInvestmentAmount float64

	investmentAmount, err := handleInput("investmentAmount", scanner)

	if err != nil {
		fmt.Println("Input failed. Please try again.")
		return
	}

	years, err = handleInput("years", scanner)

	if err != nil {
		fmt.Println("Input failed. Please try again.")
		return
	}

	if simulationType == 2 || simulationType == 3 || simulationType == 4 {
		overTimeInvestmentAmount, err = handleInput("investmentAmount2", scanner)

		if err != nil {
			fmt.Println("Input failed. Please try again.")
			return
		}
	}

	var futureValue float64

	if simulationType == 1 {
		futureValue = calculateFutureValue(investmentAmount, years)
	} else if simulationType == 2 {
		futureValue = calculateInvestmentsOverTime(52, overTimeInvestmentAmount, investmentAmount, years)
	} else if simulationType == 3 {
		futureValue = calculateInvestmentsOverTime(12, overTimeInvestmentAmount, investmentAmount, years)
	} else if simulationType == 4 {
		futureValue = calculateInvestmentsOverTime(1, overTimeInvestmentAmount, investmentAmount, years)
	}

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

func calculateInvestmentsOverTime(frequency, contributionAmount, initialInvestment, years float64) float64 {

	periodicRate := (ExpectedReturnRate / 100) / frequency
	periods := frequency * years

	initialPart := initialInvestment * math.Pow(1+periodicRate, periods)
	contributionPart := contributionAmount * ((math.Pow(1+periodicRate, periods) - 1) / periodicRate)

	return initialPart + contributionPart
}

func handleInput(version string, scanner *bufio.Scanner) (float64, error) {

	var value float64
	var err error

	for {
		switch version {
		case "investmentAmount":
			fmt.Print("Please give your investment amount: ")
		case "years":
			fmt.Print("Please give the years of the investment: ")
		case "investmentAmount2":
			fmt.Print("Please give your frequent investment amount: ")
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

// func readValue() {
//     data, err := os.ReadFile("saved_value.txt")
//     if err != nil {
//         fmt.Println("Error reading file:", err)
//         return
//     }

//     text := string(data)

//     value, err := strconv.ParseFloat(text, 64)
//     if err != nil {
//         fmt.Println("Error parsing saved value:", err)
//         return
//     }

//     fmt.Println("Loaded value:", value)
// }

// func writeValue() {
// 	value := 2500.0

//     // Convert the float to a string
//     text := strconv.FormatFloat(value, 'f', 2, 64)

//     // Save it to a file
//     err := os.WriteFile("saved_value.txt", []byte(text), 0644)
//     if err != nil {
//         fmt.Println("Error writing file:", err)
//         return
//     }

//     fmt.Println("Saved value to file.")
// }
