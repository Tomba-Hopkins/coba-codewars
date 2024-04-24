// Question :

// Your task is to sort a given string. Each word in the string will contain a single number. This number is the position the word should have in the result.
// Note: Numbers can be from 1 to 9. So 1 will be the first word (not 0).
// If the input string is empty, return an empty string. The words in the input String will only contain valid consecutive numbers.
// Examples
// "is2 Thi1s T4est 3a"  -->  "Thi1s is2 3a T4est"
// "4of Fo1r pe6ople g3ood th5e the2"  -->  "Fo1r the2 g3ood 4of th5e pe6ople"
// ""  -->  ""

// Answer :
package soal

// func Order(sentence string) string {

// 	arr := strings.Split(sentence, " ")

// 	newArr := []string{}

// for i := 0; i < len(arr); i++{
// 	for j := 0; j < len(arr); j++{
// 		str := strconv.Itoa(i + 1)
// 		if strings.Contains(arr[j], str){
// 			newArr = append(newArr, arr[j])
// 		}
// 	}
// }

// 	res := strings.Join(newArr, " ")

// 	return res
// }

// cara lain :
// func Order(sentence string) string {

// 	arr := strings.Split(sentence, " ")

// 	newArr := []string{}

// 	for i := 0; i < len(arr); i++{
// 		for j := 0; j < len(arr); j++{
// 			str := strconv.Itoa(i + 1)
// 			for k := 0; k < len(arr[j]); k++{
// 				if string(arr[j][k]) == str{
// 					newArr = append(newArr, arr[j])
// 				}
// 			}
// 		}
// 	}

// 	res := strings.Join(newArr, " ")

// 	return res
// }