package soal

import "fmt"

// GA-DE-RY-PO-LU-KI

func FindTheKey3(messages, secrets []string) (res string) {

	rumus := map[string]bool{
		"g": true,
		"a": true,
		"d": true,
		"e": true,
		"r": true,
		"y": true,
		"p": true,
		"o": true,
		"l": true,
		"u": true,
		"k": true,
		"i": true,
	}


	temp := map[string]int{}

	for i := 0; i < len(messages); i++{
		for j := 0; j < len(messages[i]); j++{
			if rumus[string(messages[i][j])]{
				karakta1 := string(messages[i][j])
				karakta2 := string(secrets[i][j])
				temp[karakta1]++
				temp[karakta2]++
			}
		}
	}


	rumus2 := map[string]int{
		"g": 1,
		"a": 1,
		"d": 2,
		"e": 2,
		"r": 3,
		"y": 3,
		"p": 4,
		"o": 4,
		"l": 5,
		"u": 5,
		"k": 6,
		"i": 6,
	}

	for _, r := range messages{
		for _, v := range r {
			for temp[string(v)] > 1 {
				temp[string(v)]--
			}
		}
	}
	for _, r := range secrets{
		for _, v := range r {
			for temp[string(v)] > 1 {
				temp[string(v)]--
			}
		}
	}
	

	urut := 1


	tempppppp := []string{}
	
	for i := 0; i < len(messages); i++{
		for j := 0; j < len(messages[i]); j++{

			text1 := string(messages[i][j])
			// text2 := string(secrets[i][j])

			if temp[text1] > 0{
				for _, v := range messages{
					for _, ko := range v {
						if  rumus2[string(ko)] == rumus2[text1] && string(ko) != text1 && rumus2[text1] == urut {
							newText := string(ko) + text1
							tempppppp = append(tempppppp, newText)
							temp[text1]--
						}
					}
				}
			} 
			urut++
		}
	}

	fmt.Println(tempppppp)
	

	return 
}

	// arr := []string{"dance on the table", "hide my beers", "scouts rocks"}
	// arr2 := []string{"egncd pn thd tgbud", "hked mr bddys", "scplts ypcis"}

	// fmt.Println(FindTheKey(arr, arr2))
	// // agdeikluopry


