// Question :

// Introduction
// The GADERYPOLUKI is a simple substitution cypher used in scouting to encrypt messages. The encryption is based on short, easy to remember key. The key is written as paired letters, which are in the cipher simple replacement.
// The most frequently used key is "GA-DE-RY-PO-LU-KI".
//  G => A
//  g => a
//  a => g
//  A => G
//  D => E
//   etc.
// The letters, which are not on the list of substitutes, stays in the encrypted text without changes.
// Other keys often used by Scouts:
// PO-LI-TY-KA-RE-NU
// KA-CE-MI-NU-TO-WY
// KO-NI-EC-MA-TU-RY
// ZA-RE-WY-BU-HO-KI
// BA-WO-LE-TY-KI-JU
// RE-GU-LA-MI-NO-WY
// Task
// Your task is to help scouts to encrypt and decrypt thier messages. Write the Encode and Decode functions.
// Input/Output
// The function should have two parameters.
// The message input string consists of lowercase and uperrcase characters and whitespace character.
// The key input string consists of only lowercase characters.
// The substitution has to be case-sensitive.
// Example
//  Encode("ABCD", "agedyropulik");             // => GBCE
//  Encode("Ala has a cat", "gaderypoluki")     // => Gug hgs g cgt
//  Decode("Dkucr pu yhr ykbir","politykarenu") // => Dance on the table
//  Decode("Hmdr nge brres","regulaminowy")  // => Hide our beers

// Answer :
package soal

import "strings"

func GaderypolukiEncodeVol2(str, key string) (res string) {

	kunci := map[string]string{}
	arr := []string{}

	for i := 0; i < len(key); i += 2{
		arr = append(arr, key[i:i+2])
	}

	for i := 0; i < len(arr); i++{
		tu := string(arr[i][0])
		wa := string(arr[i][1])
		
		kunci[tu] = wa
		kunci[wa] = tu
	}

	
	for _, r := range str{

		malesConvert := string(r)
		
		if kunci[malesConvert] != ""{
			res += kunci[malesConvert]
		} else if kunci[strings.ToLower(malesConvert)] != ""{
			norm := kunci[strings.ToLower(malesConvert)]
			cap := strings.ToUpper(norm)
			res += cap
		} else {
			res += malesConvert
		}
	}
	
	
	return 
}

func GaderypolukiDecodeVol2(str, key string) string {
	return GaderypolukiEncodeVol2(str, key)
}