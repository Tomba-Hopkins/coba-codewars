// Question :

// Introduction
// Digital Cypher assigns to each letter of the alphabet unique number. For example:
//  a  b  c  d  e  f  g  h  i  j  k  l  m
//  1  2  3  4  5  6  7  8  9 10 11 12 13
//  n  o  p  q  r  s  t  u  v  w  x  y  z
// 14 15 16 17 18 19 20 21 22 23 24 25 26
// Instead of letters in encrypted word we write the corresponding number, eg. The word scout:
//  s  c  o  u  t
// 19  3 15 21 20
// Then we add to each obtained digit consecutive digits from the key. For example. In case of key equal to 1939 :
//    s  c  o  u  t
//   19  3 15 21 20
//  + 1  9  3  9  1
//  ---------------
//   20 12 18 30 21

//    m  a  s  t  e  r  p  i  e  c  e
//   13  1 19 20  5 18 16  9  5  3  5
// +  1  9  3  9  1  9  3  9  1  9  3
//   --------------------------------
//   14 10 22 29  6 27 19 18  6  12 8
// Task
// Write a function that accepts str string and key number and returns an array of integers representing encoded str.

// Input / Output
// The str input string consists of lowercase characters only.
// The key input number is a positive integer.

// Example
// Encode("scout",1939);  ==>  [ 20, 12, 18, 30, 21]
// Encode("masterpiece",1939);  ==>  [ 14, 10, 22, 29, 6, 27, 19, 18, 6, 12, 8]

// Answer :
package soal

import "strconv"

func Encode(str string, key int) (res []int) {

	abjad := map[string]int{
		"a":1,
		"b":2,
		"c":3,
		"d":4,
		"e":5,
		"f":6,
		"g":7,
		"h":8,
		"i":9,
		"j":10,
		"k":11,
		"l":12,
		"m":13,
		"n":14,
		"o":15,
		"p":16,
		"q":17,
		"r":18,
		"s":19,
		"t":20,
		"u":21,
		"v":22,
		"w":23,
		"x":24,
		"y":25,
		"z":26,
	}

	charInt := []int{}
	for _, s := range str {
		charInt = append(charInt, abjad[string(s)])
	}


	kunci := strconv.Itoa(key)
	arrKey := []int{}
	for _, k := range kunci{
		intK, _ := strconv.Atoi(string(k))
		arrKey = append(arrKey, intK)
	}

	index := 0
	for i := 0; i < len(charInt); i++{

		if index > len(arrKey) - 1{
			index = 0
		}

		res = append(res, charInt[i] + arrKey[index])

		index++
	}

	return
}
