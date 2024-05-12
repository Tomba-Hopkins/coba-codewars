// Question :

// You will be given a list of strings. You must sort it alphabetically (case-sensitive, and based on the ASCII values of the chars) and then return the first value.
// The returned value must be a string, and have "***" between each of its letters.
// You should not remove or add elements from/to the array.

// Answer :
package soal

func TwoSort(arr []string) (res string) {


	min := arr[0]

	for i := 1; i < len(arr); i++{
		if arr[i] < min{
			min = arr[i]
		}
	}


	for j := 0; j < len(min); j++{
		if j == len(min) - 1{
			res += string(min[j])
			continue
		}

		res += string(min[j])
		for i := 0; i < 3; i++{
			res += "*"
		}
	}
	
	
	return res
}