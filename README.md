# Investment Calculator

This project is a small Go application that models investment growth and inflation-adjusted future value.

## What it does

The app lets a user:

- simulate a single investment
- simulate weekly contributions
- simulate monthly contributions
- simulate yearly contributions
- see both the nominal future value and the inflation-adjusted value

## Formula logic

The app uses a compound growth model:

- single investment: `amount * (1 + rate)^years`
- recurring contributions: the future value of a regular contribution stream is calculated using a periodic rate and contribution growth formula

## Run the app

From the project folder, run:

```bash
go run .
```

## Run the tests

```bash
go test ./...
```

## Example flow

1. Choose a simulation type from the menu
2. Enter the initial investment amount
3. Enter the number of years
4. For recurring contributions, enter the per-period contribution amount
5. Review the final value and the inflation-adjusted value

## Notes for learning

This project is a good beginner example for:

- functions
- constants and variables
- input parsing with `strconv`
- loops and menu logic
- Go testing with `testing`
- simple financial math in code

## Future ideas

Possible next upgrades include:

- saving results to a file
- loading previous values from a file
- adding monthly/yearly summaries
- letting users choose different interest rates
- adding contribution frequency validation
