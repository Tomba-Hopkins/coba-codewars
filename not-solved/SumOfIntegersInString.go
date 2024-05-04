// Question :

// Your task in this kata is to implement a function that calculates the sum of the integers inside a string. For example, in the string "The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog", the sum of the integers is 3635.
// Note: only positive integers will be tested.

// Answer :
package soal

import (
	"fmt"
	"strconv"
)

func SumOfIntegersInString(strng string) (res int) {

	a := "0123456789"
	angka := map[string]bool{}
	for _, ar := range a {
		angka[string(ar)] = true
	}

	num := []string{}
	temp := ""

	for i := 0; i < len(strng) - 1; i++{
		str := string(strng[i])

		if angka[str] && angka[string(strng[i + 1])]{
			temp += str
		} else if angka[str] {
			num = append(num, temp + str)
			temp = ""
		}
	}
	
	
	for _, numb := range num{
		intNum, _ := strconv.Atoi(numb)
		res += intNum
	}

	fmt.Println(num)
	
	return 
}