package main

import "fmt"

func mains1() {

	// 1
	println("Value 1")
	var firstName string
	firstName = "Rizky"
	println(firstName)

	// 2
	println("Value 2")
	lastName := "Kurniawan"
	println(lastName)

	// 3 combine
	println("Value 3")
	println(firstName + " " + lastName)

	// 4 OR
	fmt.Printf("halo %s %s!\n", firstName, lastName)

	// 5
	fmt.Println("halo", firstName, lastName+"!")

	// 6 Multiple variable
	fmt.Println("halo", firstName, lastName+"!")

	// 7 bisa membuat variabel seperti channel slice atau map
	newChannel := make(chan int)
	fmt.Println(newChannel)

	newSlice := []int{1, 2, 3}
	fmt.Println(newSlice)
}
