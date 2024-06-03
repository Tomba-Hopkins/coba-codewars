// Question :

// Implement a function which behaves like the uniq command in UNIX.
// It takes as input a sequence and returns a sequence in which all duplicate elements following each other have been reduced to one instance.
// Example:
// ["a", "a", "b", "b", "c", "a", "b", "c"]  =>  ["a", "b", "c", "a", "b", "c"]

// Answer :
package soal

func Uniq(a []string) (res []string) {

	if len(a) == 0 {
		return []string{}
	}

	for i := 0; i < len(a) - 1; i++{
		if a[i] != a[i + 1]{
			res = append(res, a[i])
		}
	}

	res = append(res, a[len(a) - 1])

	return 
	
}