// Question :

// Kata Task
// I have a cat and a dog.
// I got them at the same time as kitten/puppy. That was humanYears years ago.
// Return their respective ages now as [humanYears,catYears,dogYears]
// NOTES:
// humanYears >= 1
// humanYears are whole numbers only
// Cat Years
// 15 cat years for first year
// +9 cat years for second year
// +4 cat years for each year after that
// Dog Years
// 15 dog years for first year
// +9 dog years for second year
// +5 dog years for each year after that

// Answer :
package soal

func CalculateYears(years int) (result [3]int) {

	for i := 1; i <= years; i++{

		switch i{
		case 1:
			result[1] += 15
			result[2] += 15
		case 2:
			result[1] += 9
			result[2] += 9
		default:
			result[1] += 4
			result[2] += 5
		}

		result[0]++

		
	}
	
	return
}

// kalau buat var baru -> human := result[0] -> human++ malah ga berubah jir result nya