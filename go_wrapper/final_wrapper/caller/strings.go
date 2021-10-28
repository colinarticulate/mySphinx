package main

import (
	"fmt"
	"strings"
)

// Main function
func main() {

	// Creating and initializing the strings
	str1 := "Welcome, to the, online portal, of GeeksforGeeks"
	str2 := "My dog name is Dollar"
	str3 := "I like to play Ludo"

	// Displaying strings
	fmt.Println("String 1: ", str1)
	fmt.Println("String 2: ", str2)
	fmt.Println("String 3: ", str3)

	// Splitting the given strings
	// Using Split() function
	res1 := strings.Split(str1, ",")
	res2 := strings.Split(str2, "")
	res3 := strings.Split(str3, "!")
	res4 := strings.Split("", "GeeksforGeeks, geeks")

	// Displaying the result

	fmt.Println("\nResult 1: ", res1)
	fmt.Println("Result 2: ", res2)
	fmt.Println("Result 3: ", res3)
	fmt.Println("Result 4: ", res4)

	test := []string{"sil kl ay m v b sil (-3641)*word,start,end*sil,3,90*(NULL),90,90*kl,91,109*(NULL),109,109*ay,110,147*(NULL),147,147*m,148,165*(NULL),165,165*v,166,172*b,173,176*sil,177,245**0000000000000000000000000000000000000"}

	raw := strings.Split(test[0], "**")

	fmt.Printf("%T", raw)
	fields := strings.Split(raw[0], "*")

	fmt.Println(fields)
	hyp := fields[0]
	header := fields[1]

	for i := 2; i < len(fields)-1; i++ {

	}

}
