// Question :

// Complete the solution so that it returns true if the first argument(string) passed in ends with the 2nd argument (also a string).
// Examples:
// solution("abc", "bc") // returns true
// solution("abc", "d") // returns false

// Answer :
package soal

func StringEnds(str, ending string) bool {

	arr1 := []string{}
	arr2 := []string{}

	for _, s1 := range str {
		arr1 = append(arr1, string(s1))
	}
	for _, s2 := range ending {
		arr2 = append(arr2, string(s2))
	}

	temp := len(arr2) - 1


	for i := len(arr1) - 1; i >= 0; i--{
		if i == 0 && arr1[i] == arr2[temp] {
			return true
		} else if temp > 0 && arr1[i] == arr2[temp] {
			temp--
		}
	} 

	return false
	
}