package main

import (
	"math"
	"os"
	"testing"
)

func TestCalculateFutureValue(t *testing.T) {
	// We choose a small, easy-to-read input: 10 invested for 1 year.
	// This test checks that the formula returns the expected growth value.
	ratio := math.Pow(10, float64(2))
	value := math.Round(calculateFutureValue(10.0, 1.0)*ratio) / ratio

	// 10 * (1 + 5.5/100)^1 = 10.55
	if value != 10.55 {
		t.Fatalf("Value of result %f is incorrect", value)
	}
}

func TestCalculateRealFutureValue(t *testing.T) {
	// This test checks the inflation-adjusted value.
	// We use 10 invested for 10 years to make sure the adjusted value is calculated correctly.
	ratio := math.Pow(10, float64(2))
	value := math.Round(calculateRealFutureValue(10.0, 10.0)*ratio) / ratio

	// The expected result is the future value adjusted by inflation.
	if value != 7.81 {
		t.Fatalf("Value of result %f is incorrect", value)
	}
}

func TestCalculateRealFutureValueOf0(t *testing.T) {
	// Edge case: if the starting amount is 0, the result should also be 0.
	ratio := math.Pow(10, float64(2))
	value := math.Round(calculateRealFutureValue(0, 0)*ratio) / ratio

	if value != 0 {
		t.Fatalf("Value of result %f is incorrect", value)
	}
}

func TestCalculateRealFutureValueOf1000(t *testing.T) {
	// Another edge case using a larger amount and longer timeframe.
	// This helps confirm the formula still behaves correctly past the smallest examples.
	ratio := math.Pow(10, float64(2))
	value := math.Round(calculateRealFutureValue(1000, 100)*ratio) / ratio

	if value != 84.65 {
		t.Fatalf("Value of result %f is incorrect", value)
	}
}

// Positive test: valid input should be accepted and returned without an error.
func TestHandleInput_ValidAmount(t *testing.T) {
	oldStdin := os.Stdin

	// os.Pipe() creates a fake input stream that acts like a terminal.
	// r is the read end, which handleInput() will read from.
	// w is the write end, which we use to send fake keyboard input.
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	// This simulates the user typing "2500" and pressing Enter.
	_, err = w.WriteString("2500\n")
	if err != nil {
		t.Fatal(err)
	}

	// Closing the writer tells the reader that the fake input is complete.
	_ = w.Close()

	// Replace the real stdin with our fake pipe so handleInput() reads from it.
	os.Stdin = r
	defer func() {
		os.Stdin = oldStdin
	}()

	value, err := handleInput("investmentAmount")
	if err != nil {
		t.Fatalf("expected valid input to pass, got error: %v", err)
	}

	// The function should accept 2500 and return it without an error.
	if value != 2500 {
		t.Fatalf("expected value 2500, got %v", value)
	}
}

// Negative test: invalid text should not be accepted and should return an error.
func TestHandleInput_InvalidNumber(t *testing.T) {
	oldStdin := os.Stdin

	// We create another fake input stream so we can simulate bad text input.
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	// This simulates the user typing letters instead of a number.
	_, err = w.WriteString("abc\n")
	if err != nil {
		t.Fatal(err)
	}
	_ = w.Close()

	// Redirect stdin to the fake input stream.
	os.Stdin = r
	defer func() {
		os.Stdin = oldStdin
	}()

	value, err := handleInput("investmentAmount")

	// Invalid text should fail the ParseFloat step and return an error.
	if err == nil {
		t.Fatal("expected an error for invalid input")
	}

	// If the input was invalid, the function should not return a usable value.
	if value != 0 {
		t.Fatalf("expected value 0, got %v", value)
	}
}

// Negative test: a negative amount should be rejected and should not return a value.
func TestHandleInput_NegativeAmount(t *testing.T) {
	oldStdin := os.Stdin

	// This simulates the user typing -10 and pressing Enter.
	r, w, _ := os.Pipe()
	_, _ = w.WriteString("-10\n")
	_ = w.Close()

	// Redirect stdin to the fake input stream.
	os.Stdin = r
	defer func() {
		os.Stdin = oldStdin
	}()

	value, err := handleInput("investmentAmount")

	// Negative values should be rejected by validation.
	if err == nil {
		t.Fatal("expected an error for negative amount")
	}

	// The function should not silently continue with a valid numeric value.
	if value != 0 {
		t.Fatalf("expected value 0, got %v", value)
	}
}
