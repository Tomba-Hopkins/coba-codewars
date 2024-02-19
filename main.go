package main

import "fmt"

func HowMuchILoveYou(i int) string {
	arr := []string{"I love you", "a little", "a lot", "passionately", "madly", "not at all"}
	
	if i > len(arr) {
		return arr[(i - 1) % len(arr)]
	} else {
		return arr[i - 1]
	}
}

func main() {
	fmt.Println(HowMuchILoveYou(255))
}