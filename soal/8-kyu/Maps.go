// Question :
// Given an array of integers, return a new array with each value doubled.
// For example:
// [1, 2, 3] --> [2, 4, 6]

// Answer :
package soal

func Maps(x []int) []int {
	result := []int{}

	for _, xixixi := range x {
		xixixi *= 2
		result = append(result, xixixi)
	}
	return result
}