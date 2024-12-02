package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func bmi_calculator(weight, height float64) (bmi float64) {
	bmi = weight / math.Pow(height, 2)
	bmi, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", bmi), 64)
	return
}

func normalizeInputString(inputNum string) (parsedStr string) {
	parsedStr = strings.ReplaceAll(inputNum, "\n", "")
	return
}

func convertInputNumbersToFloat64(inputNum string) (convertedNum float64) {
	convertedNum, _ = strconv.ParseFloat(inputNum, 32)
	return
}

func main() {
	// Get User inputs:
	fmt.Println("BMI Calculator:")
	fmt.Println("----------------")
	fmt.Print("Please enter your weight (kg): ")
	weightInput, _ := reader.ReadString('\n')
	weightInput = normalizeInputString(weightInput)
	fmt.Print("Please enter your height (m): ")
	heightInput, _ := reader.ReadString('\n')
	heightInput = normalizeInputString(heightInput)

	// Parse Input values and convert them into float numbers
	weight := convertInputNumbersToFloat64(weightInput)
	height := convertInputNumbersToFloat64(heightInput)

	// Calculate BMI
	bmi := bmi_calculator(weight, height)
	fmt.Println(bmi)
}
