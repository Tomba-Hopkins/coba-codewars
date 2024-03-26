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
// Write a function that accepts an array of integers code and a key number. As the result, it should return string containg a decoded message from the code.
// Input / Output
// The code is a array of positive integers.
// The key input is a nonnegative integer.
// Example
// decode([ 20, 12, 18, 30, 21],1939);  ==> "scout"
// decode([ 14, 10, 22, 29, 6, 27, 19, 18, 6, 12, 8],1939);

// Answer :
package soal

import "strconv"

func DecodeVol2(code []int, key int) string {

	res := ""
	abjad := map[int]string{}

	forMap := "abcdefghijklmnopqrstuvwxyz"
	for f := 0; f < len(forMap); f++ {
		abjad[f+1] = string(forMap[f])
	}

	strKey := strconv.Itoa(key) 

	newdex := 0
	for i := 0; i < len(code); i++{
		if newdex == len(strKey) {
			newdex = 0
		}
		temp, _ := strconv.Atoi(string(strKey[newdex]))
		res += abjad[code[i] - temp]
		newdex++
	}

	return res
}

// yang asli pake js go nya ga ada : 
// function decode  (code, n) {
//     let result = ""
//     let buatMap = "abcdefghijklmnopqrstuvwxyz"
//     let abjad = new Map()
//     for (let i = 0; i < buatMap.length; i++){
//         abjad.set(i + 1,buatMap[i])
//     }
//     let NewN = n.toString()
//     let temp = 0
//     for (let i = 0; i < code.length; i++){
//         if (temp == NewN.length) {
//             temp = 0
//         }
//         result += abjad.get(code[i] - parseInt(NewN[temp]))
//         temp++
//     }
//     return result
// }