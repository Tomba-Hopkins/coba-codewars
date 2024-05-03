// Question :

// Complete the solution so that it splits the string into pairs of two characters. If the string contains an odd number of characters then it should replace the missing second character of the final pair with an underscore ('_').
// Examples:
// * 'abc' =>  ['ab', 'c_']
// * 'abcdef' => ['ab', 'cd', 'ef']

// Answer :
package soal

func SplitStrings(str string) (res []string) {

	if len(str) % 2 == 0 {
		for i := 0; i < len(str); i += 2{
			res = append(res, string(str[i])+string(str[i + 1]))
		}
	} else {
		for i := 0; i < len(str) + 1; i += 2{
			if i == len(str) - 1{
				res = append(res, string(str[i])+"_")
			} else {
				res = append(res, string(str[i]) + string(str[i + 1]))
			}
		}
	}

	return

}