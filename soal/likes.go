// Question :
// You probably know the "like" system from Facebook and other pages. People can "like" blog posts, pictures or other items. We want to create the text that should be displayed next to such an item.
// Implement the function which takes an array containing the names of people that like an item. It must return the display text as shown in the examples:
// []                                -->  "no one likes this"
// ["Peter"]                         -->  "Peter likes this"
// ["Jacob", "Alex"]                 -->  "Jacob and Alex like this"
// ["Max", "John", "Mark"]           -->  "Max, John and Mark like this"
// ["Alex", "Jacob", "Mark", "Max"]  -->  "Alex, Jacob and 2 others like this"
// Note: For 4 or more names, the number in "and 2 others" simply increases.

// Answer :
package soal

import "strconv"
func likes(people []string) string {
	result := ""
	if len(people) == 0 {
		result = "no one likes this"
	} else if len(people) == 1 {
		result = people[0] + " likes this"
	} else if len(people) == 2 {
		result = people[0] + " and " + people[1] + " like this"
	} else if len(people) == 3 {
		result = people[0] + ", " + people[1] + " and " + people[2] + " like this"
	} else {
		result = people[0] + ", " + people[1] + " and " + strconv.Itoa(len(people) - 2) + " others like this"
	
	}
	return result
}

// js :
// function likes(names) {
// 	let result = '';
// 	  if (names.length === 0) {
// 	   result = "no one likes this";
// 	  } else if (names.length === 1) {
// 	   result =  names[0] + " likes this";
// 	  } else if (names.length === 2) {
// 	   result =  names[0] + " and " + names[1] + " like this";
// 	  } else if (names.length === 3) {
// 	   result = names[0] + ", " + names[1] + " and " + names[2] + " like this";
// 	  } else {
// 		result = names[0] + ", " + names[1] + " and " + (names.length - 2) + " others like this";
// 	  }
// 	  return result;
// 	}