// Question :

// Write an algorithm that takes an array and moves all of the zeros to the end, preserving the order of the other elements.
// MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}) // returns []int{ 1, 2, 1, 1, 3, 1, 0, 0, 0, 0 }

// Answer :
package soal

func MoveZeros(arr []int) (res []int) {

	z := 0

	for _, r := range arr {
		if r == 0 {
			z++
			continue
		}
		res = append(res, r)
	}

	for i := 0; i < z; i++ {
		res = append(res, 0)
	}

	return
}