package main

import "fmt"

func main() {
	var myName = "Kresna"
	fmt.Println("My name is", myName, myName)

	var name string = "Bagus"
	fmt.Println("My son is", name)

	userName := "admin"
	fmt.Println("username =", userName)

	var sum int 
	// the default is 0
	fmt.Println("The sum is", sum)

	// compound
	part1, other := 1, 5
	fmt.Println("part1 =", part1, "other =", other)

	part2, other := 2, 0
	fmt.Println("part2 =", part2, "other =", other)

	sum = part1 + part2
	fmt.Println("sum =", sum)

	// block assignment
	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println("Lesson name =", lessonName)
	fmt.Println("lessonType =", lessonType)

	// Compound assignment where ignore one of the variables
	word1, word2, _ := "hello", "world", "!"
	fmt.Println(word1, word2)

}
