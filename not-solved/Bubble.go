// Question :

// #Bubbleing around
// Since everybody hates chaos and loves sorted lists we should implement some more sorting algorithms. Your task is to implement a Bubble sort (for some help look at https://en.wikipedia.org/wiki/Bubble_sort) and return a list of snapshots after each change of the initial list.
// e.g.
// If the initial list would be l=[1,2,4,3] my algorithm rotates l[2] and l[3] and after that it adds [1,2,3,4] to the result, which is a list of snapshots.
// [1,2,4,3] should return [ [1,2,3,4] ]
// [2,1,4,3] should return [ [1,2,4,3], [1,2,3,4] ]
// [1,2,3,4] should return []

// Answer :
package soal

func Bubble(array []int) [][]int {

	result := [][]int{}

	for i := 0; i < len(array) - 1; i++{
		swap := false

		for j := 0; j < len(array) - i - 1; j++{
			if array[j] > array[j + 1] {
				array[j], array[j + 1] = array[j + 1], array[j]
				result = append(result, make([]int, len(array)))
				copy(result[i], array)
				result = append(result, result[i])
				swap = true
			}
		}

		if !swap{
			break
		}
		

	}
	return result
	
}