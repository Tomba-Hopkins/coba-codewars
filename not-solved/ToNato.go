// Question :
// Task
// You'll have to translate a string to Pilot's alphabet (NATO phonetic alphabet).
// Input:
// If, you can read?
// Output:
// India Foxtrot , Yankee Oscar Uniform Charlie Alfa November Romeo Echo Alfa Delta ?
// Note:
// There is a preloaded dictionary that you can use, named NATO. It uses uppercase keys, e.g. NATO['A'] is "Alfa". See comments in the initial code to see how to access it in your language.
// The set of used punctuation is ,.!?.
// Punctuation should be kept in your return string, but spaces should not.
// Xray should not have a dash within.
// Every word and punctuation mark should be seperated by a space ' '.
// There should be no trailing whitespace

// Answer :
package soal

import "strings"

func ToNato(words string) (res string) {

	kamus := map[string]string{}
	kamusVal := []string{"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November", "Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor", "Whiskey", "Xray", "Yankee", "Zulu"}

	for i := 0; i < len(kamusVal); i++ {

		char := string(kamusVal[i][0])
		kamus[char] = kamusVal[i]

	}

	for _, w := range words{
		char := strings.ToUpper(string(w))

		if kamus[char] != ""{
			res += kamus[char] + " "
		}else if char == " "{
			continue
		} else {
			res += string(w)
		}
	}

	if res[len(res) - 1] == ' ' && len(res) > 1{
		res = res[:len(res) - 1 ]
	}

	return 
}