// Question :

// Write a function that accepts an array of 10 integers (between 0 and 9), that returns a string of those numbers in the form of a phone number.
// Example
// CreatePhoneNumber([10]uint{1,2,3,4,5,6,7,8,9,0})  // returns "(123) 456-7890"
// The returned format must be correct in order to complete this challenge.
// Don't forget the space after the closing parentheses!

// Answer :
package soal

import "strconv"


func CreatePhoneNumber(numbers [10]uint) (res string) {
	for i := 0; i < len(numbers); i++{
		if i == 0 {
			res += "("
		}
		
		res += strconv.Itoa(int(numbers[i]))

		if i == 2 {
			res += ")" + " "
		} else if i == 5 {
			res += "-"
		}
	}
	return
}