// Question :
// Create a function that returns the CSV representation of a two-dimensional numeric array.
// Example:
// input:
//    [[ 0, 1, 2, 3, 4 ],
//     [ 10,11,12,13,14 ],
//     [ 20,21,22,23,24 ],
//     [ 30,31,32,33,34 ]]
// output:
//      '0,1,2,3,4\n'
//     +'10,11,12,13,14\n'
//     +'20,21,22,23,24\n'
//     +'30,31,32,33,34'
// Array's length > 2.
// More details here: https://en.wikipedia.org/wiki/Comma-separated_values
// Note: you shouldn't escape the \n, it should work as a new line.

// Answer :
package soal

import "strconv"

func ToCsvText(array [][]int) string {
	result := ""
	for i := 0; i < len(array); i++{
		for j := 0; j < len(array[i]); j++{
			text := strconv.Itoa(array[i][j])
			result += text
			if j < len(array[i]) - 1 {
				 result +=  ","
			}
		}
		if i < len(array) {
			result += "\n"
		}
	}
	return result
}