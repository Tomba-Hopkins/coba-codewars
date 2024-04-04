// Question :

// Write a function that takes an integer as input, and returns the number of bits that are equal to one in the binary representation of that number. You can guarantee that input is non-negative.
// Example: The binary representation of 1234 is 10011010010, so the function should return 5 in this case

// Answer :
package soal

import "strconv"


func CountBits(input uint) (res int) {

	binary := strconv.FormatInt(int64(input), 2)
	
	for _, b := range binary {
		if string(b) == "1"{
			res++
		}
	}

	return

}