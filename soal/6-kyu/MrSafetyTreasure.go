// Question :

// Introduction
// Mr. Safety loves numeric locks and his Nokia 3310. He locked almost everything in his house. He is so smart and he doesn't need to remember the combinations. He has an algorithm to generate new passcodes on his Nokia cell phone.
// Task
// Can you crack his numeric locks? Mr. Safety's treasures wait for you. Write an algorithm to open his numeric locks. Can you do it without his Nokia 3310?
// Input
// The str or message (Python) input string consists of lowercase and upercase characters. It's a real object that you want to unlock.
// Output
// Return a string that only consists of digits.
// Example
// ``` unlock("Nokia") // => 66542 unlock("Valut") // => 82588 unlock("toilet") // => 864538 ```

// Answer :
package soal

import "strings"

func MrSafetyTreasure(str string) (res string) {

	abjad := map[string]string{}

	determine := '2'
	strArr := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	for i := 0; i < len(strArr); i++{
		for j := 0; j < len(strArr[i]); j++{
			abjad[string(strArr[i][j])] = string(determine)
		}
		determine++
	}

	for i := 0; i < len(str); i++{
		key := strings.ToLower(string(str[i]))
		res += abjad[key]
	}
	
	return
}