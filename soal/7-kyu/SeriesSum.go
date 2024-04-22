// Question :

// Task
// Your task is to write a function which returns the n-th term of the following series, which is the sum of the first n terms of the sequence (n is the input parameter).
// You will need to figure out the rule of the series to complete this.

// Rules
// You need to round the answer to 2 decimal places and return it as String.

// If the given value is 0 then it should return "0.00".

// You will only be given Natural Numbers as arguments.

// Examples (Input --> Output)
// n
// 1 --> 1 --> "1.00"
// 2 --> 1 + 1/4 --> "1.25"
// 5 --> 1 + 1/4 + 1/7 + 1/10 + 1/13 --> "1.57"

// Answer :
package soal

import "strconv"

func SeriesSum(n int) string {

	if n < 1 {
		return "0.00"
	}

	one := 1.00
	per := 4.00

	for i := 1; i < n; i++ {
		one += 1.00 / per
		per += 3.00
	}

	// res := fmt.Sprintf("%.2f", one)

	res := strconv.FormatFloat(one, 'f', 2, 64)

	return res
}
