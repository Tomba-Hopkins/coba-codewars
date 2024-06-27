// Question :

// Determine the total number of digits in the integer (n>=0) given as input to the function. For example, 9 is a single digit, 66 has 2 digits and 128685 has 6 digits. Be careful to avoid overflows/underflows.
// All inputs will be valid.

// Answer :
package soal

import "strconv"

func NumberOfDecimalDigit(n uint64) int {

	strLah := strconv.FormatUint(n, 10)
	return len(strLah)
	
}