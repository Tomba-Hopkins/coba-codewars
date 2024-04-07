// Question :

// You are given an array (which will have a length of at least 3, but could be very large) containing integers. The array is either entirely comprised of odd integers or entirely comprised of even integers except for a single integer N. Write a method that takes the array as an argument and returns this "outlier" N.
// Examples
// [2, 4, 0, 100, 4, 11, 2602, 36] -->  11 (the only odd number)
// [160, 3, 1719, 19, 11, 13, -21] --> 160 (the only even number)

// Answer :
package soal

func FindOutlier(integers []int) int {

	odd := []int{}
	even := []int{}

	for _, i := range integers {

		if i % 2 == 0 {
			even = append(even, i)
		} else {
			odd = append(odd, i)
		}
	}

	if len(odd) > len(even) {
		return even[0]
	} else if len(even) > len(odd) {
		return odd[0]
	} else {
		return integers[0]
	}

}