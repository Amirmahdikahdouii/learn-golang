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

func bmi_calculator(weight, height float64) float64 {
	return weight / math.Pow(height, 2)
}

func main() {
	// Get User inputs:
	fmt.Println("BMI Calculator:")
	fmt.Println("----------------")
	fmt.Print("Please enter your weight (kg): ")
	weightInput, _ := reader.ReadString('\n')
	weightInput = strings.ReplaceAll(weightInput, "\n", "")
	fmt.Print("Please enter your height (m): ")
	heightInput, _ := reader.ReadString('\n')
	heightInput = strings.ReplaceAll(heightInput, "\n", "")

	// Parse Input values and convert them into float numbers
	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	// Calculate BMI
	bmi := bmi_calculator(weight, height)
	fmt.Println(bmi)
}
