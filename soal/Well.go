// Question :

// For every good kata idea there seem to be quite a few bad ones!
// In this kata you need to check the provided array (x) for good ideas 'good' and bad ideas 'bad'. If there are one or two good ideas, return 'Publish!', if there are more than 2 return 'I smell a series!'. If there are no good ideas, as is often the case, return 'Fail!'.

// Answer :
package soal

func Well(x []string) string {
	res := 0
	for _, r := range x {
		if r == "good" {
			res++
		}
	}
	if res <= 2 && res != 0{
		return "Publish!"
	} else if res > 2 {
		return "I smell a series!"
	} 
	return "Fail!"

}