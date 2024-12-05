package main

import "fmt"

func mains2() {

	// 1
	var decimal float64 = 3.14
	fmt.Println(decimal)

	// 2
	var isTrue bool = true
	fmt.Println(isTrue)

	// 3
	var name string = "Rizky"
	fmt.Println(name)

	// 4
	var age int = 20
	fmt.Println(age)

	// 5
	var (
		firstName = "Rizky"
		lastName  = "Kurniawan"
	)
	fmt.Println(firstName, lastName)

	// 6

	var nominal uint64 = 1000000
	fmt.Println(nominal)

	// konstanta

	const dataPI = 3.14
	fmt.Println(dataPI)

	// multiple declaration constant
	const (
		square   string  = "kotak"
		isToday  bool    = true
		numeric  uint8   = 1
		floatNum float64 = 2.2
	)

	fmt.Println(square, isToday, numeric, floatNum)

}
