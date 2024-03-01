// Question :
// Move the first letter of each word to the end of it, then add "ay" to the end of the word. Leave punctuation marks untouched.
// Examples
// pigIt('Pig latin is cool'); // igPay atinlay siay oolcay
// pigIt('Hello world !');     // elloHay orldway !

// Answer :
package soal

import "fmt"

func PigLatin(text string) string {
	result := ""
	arr := []string{}
	for i := 0; i < len(text); i++{
		str := string(text[i])

		if str == " "{
			arr = append(arr, text[:i])
			arr = append(arr, text[i + 1:]) // awalnya text[i:] doang terus salah
		}
	}
	fmt.Println(arr)
	fmt.Println(len(arr))

	for i := 0; i < len(arr); i++{
		for j := 0; j < len(arr[i]); j++{
			if j != 0 {
				result += string(arr[i][j])
			}
		}
		result += string(arr[i][0]) + "ay" + " "
	}

	fmt.Println(result)
	return ""
}

// func PigLatin(text string) string {
// 	result := ""

// 	arr := strings.Split(text, " ")
// 	fmt.Println(arr)
// 	fmt.Println(len(arr))

// 	return result
// }