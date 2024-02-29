// Question :
// Complete the solution so that the function will break up camel casing, using a space between words.
// Example
// "camelCasing"  =>  "camel Casing"
// "identifier"   =>  "identifier"
// ""             =>  ""

// Answer :
package soal

import "strings"

func Solution(text string) string {
	result := ""
	for i := 0; i < len(text); i++ {
		stringText := string(text[i])
		if stringText == strings.ToUpper(stringText) {
			result += " "
		}
		result += stringText
	}
	return result
}

// js :
// function solution(string) {
// 	let result = ""
// 	for(let i = 0; i < string.length; i++){
// 		if(string[i] === string[i].toUpperCase()){
// 		  result += " "
// 		}
// 		result += string[i]
// 	}
// 	return result
//   }