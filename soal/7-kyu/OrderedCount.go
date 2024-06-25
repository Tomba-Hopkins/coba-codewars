// Question :

// Count the number of occurrences of each character and return it as a (list of tuples) in order of appearance. For empty output return (an empty list).
// Consult the solution set-up for the exact data structure implementation depending on your language.
// Example:
// OrderedCount("abracadabra") == []Tuple{Tuple{'a', 5}, Tuple{'b', 2}, Tuple{'r', 2}, Tuple{'c', 1}, Tuple{'d', 1}}
// // Where
// type Tuple struct {
//     Char  rune
//     Count int
// }

// Answer :
package soal

type Tuple struct {
	Char string
	Num  int
}

func OrderedCount(text string) []Tuple {

	Res := []Tuple{}
	maap := map[rune]int{}
	newMaap := map[rune]int{}

	
	for _, t := range text {
		maap[t] = 1
		newMaap[t]++
	}

	newText := ""

	for _, t2 := range text{
		if maap[t2] > 0 {
			newText += string(t2)
			maap[t2]--
		}
	}


	for _, r := range newText {
		Res = append(Res, Tuple{string(r), newMaap[r]})
	}

	return Res

}