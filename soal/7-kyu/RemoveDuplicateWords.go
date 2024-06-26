// Question :

// Your task is to remove all duplicate words from a string, leaving only single (first) words entries.
// Example:
// Input:
// 'alpha beta beta gamma gamma gamma delta alpha beta beta gamma gamma gamma delta'
// Output:
// 'alpha beta gamma delta'

// Answer :
package soal

import "strings"

// Ini salah gatau kenapa padahal mah udah bener outputnya
func RemoveDuplicateWords(str string) (res string) {

	arr := strings.Split(str, " ")
	duplikatGa := map[string]int{}

	for _, a := range arr {
		duplikatGa[a] = 1
	}

	for _, n := range arr{
		if duplikatGa[n] > 0 {
			res += n + " "
			duplikatGa[n]--
		}
	}

	return res[:len(res) - 1]
}

// ini versi js nya malah bener
// function removeDuplicateWords(str) {
//     const arr = str.split(" ");
//     const duplikatGa = {};
//     arr.forEach(a => {
//         duplikatGa[a] = 1;
//     });
//     let res = "";
//     arr.forEach(n => {
//         if (duplikatGa[n] > 0) {
//             res += n + " ";
//             duplikatGa[n]--;
//         }
//     });
//     return res.trim();
// }
