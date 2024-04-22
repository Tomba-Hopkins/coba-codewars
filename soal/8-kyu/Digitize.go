// Question :

// Convert number to reversed array of digits
// Given a random non-negative number, you have to return the digits of this number within an array in reverse order.
// Example(Input => Output):
// 35231 => [1,3,2,5,3]
// 0 => [0]

// Answer :
package soal

import "strconv"

func Digitize(n int) (res []int) {

	new := strconv.Itoa(n)

	for i := len(new) - 1; i >= 0; i--{
		newInt, _ := strconv.Atoi(string(new[i]))
		res = append(res, newInt)
	}

	return
}