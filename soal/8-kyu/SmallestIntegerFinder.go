// Question :

// For example:
// Given [34, 15, 88, 2] your solution will return 2
// Given [34, -345, -1, 100] your solution will return -345
// You can assume, for the purpose of this kata, that the supplied array will not be empty.

// Answer :
package soal

func SmallestIntegerFinder(numbers []int) (res int) {

	res = numbers[0]

	for i := 1; i < len(numbers); i++{
		if numbers[i] < res {
			res = numbers[i]
		}
	}
	
	return
}