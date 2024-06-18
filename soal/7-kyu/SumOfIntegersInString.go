// Question :

// Your task in this kata is to implement a function that calculates the sum of the integers inside a string. For example, in the string "The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog", the sum of the integers is 3635.
// Note: only positive integers will be tested.

// Answer :
package soal

import (
	"strconv"
)

func SumOfIntegersInString(strng string) (res int) {

	a := "0123456789"
	angka := map[string]bool{}
	for _, ar := range a {
		angka[string(ar)] = true
	}

	temp := ""
	strArr := []string{}

	for i := 0; i < len(strng); i++{
		if angka[string(strng[i])]{
			temp += string(strng[i])
		} else if temp != "" {
			strArr = append(strArr, temp)
			temp = ""
		}
		
	}

	if temp != ""{
		strArr = append(strArr, temp)
	}

	for _, i := range strArr{
		intI, _ := strconv.Atoi(i)
		res += intI
	}

	return
}