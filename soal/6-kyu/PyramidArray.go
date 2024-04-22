// Question :

// Write a function that when given a number >= 0, returns an Array of ascending length subarrays.
// pyramid(0) => [ ]
// pyramid(1) => [ [1] ]
// pyramid(2) => [ [1], [1, 1] ]
// pyramid(3) => [ [1], [1, 1], [1, 1, 1] ]
// Note: the subarrays should be filled with 1s

// Answer :
package soal

func Pyramid(n int) [][]int {

	res := [][]int{}

	for i := 1; i <= n; i++ {
		res = append(res, make([]int, i))
		
		temp := 0
		for temp < i {
			res[i - 1][temp] = 1
			temp++
		}
	}

	return res
}