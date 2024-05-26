// Question :

// The two oldest ages function/method needs to be completed. It should take an array of numbers as its argument and return the two highest numbers within the array. The returned value should be an array in the format [second oldest age,  oldest age].
// The order of the numbers passed in could be any order. The array will always include at least 2 items. If there are two or more oldest age, then return both of them in array format.
// For example (Input --> Output):
// [1, 2, 10, 8] --> [8, 10]
// [1, 5, 87, 45, 8, 8] --> [45, 87]
// [1, 3, 10, 0]) --> [3, 10]

// Answer :
package soal

func TwoOldestAges(ages []int) (res [2]int) {

	secMax := 0
	max := 0

	val := map[int]int{}

	for _, r := range ages{
		val[r]++
	}

	for _, f := range ages{
		if f > secMax && val[f] > 0{
			secMax = f
		}
	}
	val[secMax]--

	for _, s := range ages{
		if s > max && val[s] > 0{
			max = s
		}
	}
	val[max]--

	res[0] = max
	res[1] = secMax

	return
	

}