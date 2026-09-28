package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func TestCalculateFutureValue(t *testing.T) {
	got := calculateFutureValue(1000, 10)
	want := 1708.14

	if math.Abs(got-want) > 0.01 {
		t.Fatalf("calculateFutureValue(1000, 10) = %f; want %f", got, want)
	}
}

func TestCalculateRealFutureValue(t *testing.T) {
	got := calculateRealFutureValue(1708.14, 10)
	want := 1334.4

	if math.Abs(got-want) > 0.1 {
		t.Fatalf("calculateRealFutureValue(1708.14, 10) = %f; want %f", got, want)
	}
}

func TestCalculateInvestmentsOverTimeWeekly(t *testing.T) {
	got := calculateInvestmentsOverTime(52, 10, 1000, 10)
	want := 8660.56

	if math.Abs(got-want) > 0.01 {
		t.Fatalf("weekly contributions = %f; want %f", got, want)
	}
}

func TestCalculateInvestmentsOverTimeMonthly(t *testing.T) {
	got := calculateInvestmentsOverTime(12, 10, 1000, 10)
	want := 3326.15

	if math.Abs(got-want) > 0.01 {
		t.Fatalf("monthly contributions = %f; want %f", got, want)
	}
}

func TestCalculateInvestmentsOverTimeYearly(t *testing.T) {
	got := calculateInvestmentsOverTime(1, 10, 1000, 10)
	want := 1836.9

	if math.Abs(got-want) > 0.01 {
		t.Fatalf("yearly contributions = %f; want %f", got, want)
	}
}

func TestHandleInputValidAmount(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader("2500\n"))

	got, err := handleInput("investmentAmount", scanner)
	if err != nil {
		t.Fatalf("expected valid input to succeed, got error: %v", err)
	}

	if got != 2500 {
		t.Fatalf("handleInput returned %f; want 2500", got)
	}
}

func TestHandleInputInvalidNumber(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader("abc\n"))

	got, err := handleInput("investmentAmount", scanner)
	if err == nil {
		t.Fatal("expected invalid input to fail")
	}

	if got != 0 {
		t.Fatalf("expected 0 value on invalid input, got %f", got)
	}
}

func TestHandleInputNegativeValue(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader("-10\n"))

	got, err := handleInput("investmentAmount", scanner)
	if err == nil {
		t.Fatal("expected negative value to fail")
	}

	if got != 0 {
		t.Fatalf("expected 0 value for invalid negative input, got %f", got)
	}
}
