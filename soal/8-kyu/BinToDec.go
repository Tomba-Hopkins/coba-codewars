// Question :

// Complete the function which converts a binary number (given as a string) to a decimal number.

// Answer :

package soal

import "strconv"

func BinToDec(bin string) int {
	res, _ :=  strconv.ParseInt(bin, 2, 12)
	return int(res)
}