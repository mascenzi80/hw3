package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

// Function to generate a binary string of a given size
func binary_generator(size int) string {

	// find max value,
	// generate a random number between 0 and max value
	// format the number to binary string
	maxValue := int(math.Pow(2, float64(size)))
	num := rand.Intn(maxValue)
	format := fmt.Sprintf("%%0%db", size)
	return fmt.Sprintf(format, num)

}

// Function to add two binary strings
func add_multiplicand_2_product(product, multiplicand string) string {
	result := ""
	carry := 0

	// iterate through the binary strings
	for idx := len(product) - 1; idx >= 0; idx-- {
		// assign bits at index of both strings
		// convert from char to int
		product_bit := int(product[idx] - '0')
		multiplicand_bit := int(multiplicand[idx] - '0')

		// add both bits and carry
		sum := product_bit + multiplicand_bit + carry

		// calculate new carry and result
		carry = sum / 2
		result = string(sum%2+'0') + result
	}

	// add carry to result if carry is greater than 0
	if carry > 0 {
		result = "1" + result
	}

	return result
}

// Function for shifting multiplicand left
func shift_multiplicand_left(multiplicand string) string {
	return multiplicand[1:] + "0"
}

// Function for shifting multiplier right
func shift_multiplier_right(multiplier string) string {
	return "0" + multiplier[:len(multiplier)-1]
}

// Function to pad a binary integer string
// to twice the size of the original bit size
func pad(num string, size int) string {
	return strings.Repeat("0", size) + num
}

func main() {
	// Select Binary size
	bitSize := 32
	product := strings.Repeat("0", bitSize*2)

	// Generate two random binary integer strings
	num_1 := binary_generator(bitSize)
	num_2 := binary_generator(bitSize)

	// assign the multiplier and multiplicand
	// set multiplicand to double size of binary size
	multiplier := num_1
	multiplicand := pad(num_2, bitSize)

	// iterate through the multiplier
	for i := 0; i < len(multiplier); i++ {

		// test lsb of multiplier
		if multiplier[len(multiplier)-1] == '1' {
			// if lsb is 1, add multiplicand to product
			product = add_multiplicand_2_product(product, multiplicand)
		}

		// shift multiplicand left and multiplier right
		multiplicand = shift_multiplicand_left(multiplicand)
		multiplier = shift_multiplier_right(multiplier)
	}

	// trim leading '0'
	product = strings.TrimLeft(product, "0")

	padding := len(product) - (len(num_2))

	fmt.Printf(
		"%s%s\n\033[4m%s* %s\033[0m\n%s\n",
		strings.Repeat(" ", padding),
		num_1,
		strings.Repeat(" ", padding-2),
		num_2,
		product)
}
