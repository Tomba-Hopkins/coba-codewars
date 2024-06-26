// Question :

// In this little assignment you are given a string of space separated numbers, and have to return the highest and lowest number.

// Examples
// HighAndLow("1 2 3 4 5")  // return "5 1"
// HighAndLow("1 2 -3 4 5") // return "5 -3"
// HighAndLow("1 9 3 4 -5") // return "9 -5"
// Notes
// All numbers are valid Int32, no need to validate them.
// There will always be at least one number in the input string.
// Output string must be two numbers separated by a single space, and highest number is first.

// Answer :
package soal

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	arr := strings.Split(in, " ")

	high, _ := strconv.Atoi(arr[0])
	low, _ := strconv.Atoi(arr[0])

	
	for i := 0; i < len(arr); i++{
		temp, _ := strconv.Atoi(arr[i])

		if temp > high{
			high = temp
		}

		if temp < low {
			low = temp
		}
	}

	 return fmt.Sprintf("%d %d", high, low)
	
}