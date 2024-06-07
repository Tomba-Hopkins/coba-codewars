// Question :

// Build Tower
// Build a pyramid-shaped tower, as an array/list of strings, given a positive integer number of floors. A tower block is represented with "*" character.
// For example, a tower with 3 floors looks like this:
// [
//   "  *  ",
//   " *** ",
//   "*****"
// ]
// And a tower with 6 floors looks like this:
// [
//   "     *     ",
//   "    ***    ",
//   "   *****   ",
//   "  *******  ",
//   " ********* ",
//   "***********"
// ]

// Answer :
package soal

func TowerBuilder(nFloors int) (res []string) {

	for i := 0; i < nFloors; i++{

		session := ""
		space := " "
		star := "*"


		for j := 0; j < nFloors - i - 1; j++{
			session += space
		}
		
		for j := 0; j < 2 * i + 1; j++{
			session += star
			}
			
		for j := 0; j < nFloors - i - 1; j++{
			session += space
		}
	

		res = append(res, session)
		
	}
	
	return
}