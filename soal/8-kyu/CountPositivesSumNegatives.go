// Question :

// Given an array of integers.
// Return an array, where the first element is the count of positives numbers and the second element is sum of negative numbers. 0 is neither positive nor negative.
// If the input is an empty array or is null, return an empty array.
// Example
// For input [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15], you should return [10, -65].

// Answer :
package soal


func CountPositivesSumNegatives(numbers []int) (res []int) {
	pos := 0
	sum := 0
	for _, n := range numbers {
		if n > 0 {
			pos++
		}
		if n < 0 {
			sum += n
		}
	}
	res = append(res, pos)
	res = append(res, sum)
	return res
}