// Question :

// Given a string of digits, you should replace any digit below 5 with '0' and any digit 5 and above with '1'. Return the resulting string.
// Note: input will never be an empty string

// Answer :
package soal

import "strconv"

func FakeBin(x string) (res string) {

	arr := []string{}

	for _, r := range x {
		arr = append(arr, string(r))
	}

	for _, y := range arr {
		num, _ := strconv.Atoi(y)

		if num >= 5 {
			res += "1"
		} else {
			res += "0"
		}
	}

	return

}