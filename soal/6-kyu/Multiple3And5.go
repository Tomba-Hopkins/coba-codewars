// Question :
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
// Finish the solution so that it returns the sum of all the multiples of 3 or 5 below the number passed in.
// Note: If the number is a multiple of both 3 and 5, only count it once.

// Answer :
package soal

func Multiple3And5(number int) int {

	temp := 0

	for i := 0; i < number; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			temp += i
		}
	}

	return temp

}
