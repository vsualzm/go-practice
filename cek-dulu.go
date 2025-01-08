package main

import (
	"fmt"
	"time"
)

// go practice
// func main() {

// Bermain dengna go routien
// dengan function PrintMessage
// case 1 // coba jalankan
// go PrintMessage("Hello from Goroutine!") // Goroutine
// PrintMessage("Hello from Main!")         // Fungsi utama

// case 2

// berapa core yang akan digunakan
// runtime.GOMAXPROCS(2)

// ini adalah go routine di awali dengan go
// go Print(5, "Hello from goroutine!")
// Print(5, "Hello from main!")

// var input string
// fmt.Scanln(&input)

// case 3
// Golang Channel

// runtime.GOMAXPROCS(2)

// // buat channel
// var messages = make(chan string)

// var sayHelloTo = func(who string) {
// 	var data = fmt.Sprintf("Hello %s", who)
// 	// kirim data ke channel
// 	messages <- data
// }

// go sayHelloTo("Eko")
// go sayHelloTo("Joko")
// go sayHelloTo("Budi")

// nah ketika di sini hasil nya akan berbeda saat di print
// di karena go routine dijalankan secara paralel

// var message1 = <-messages
// fmt.Println(message1)
// var message2 = <-messages
// fmt.Println(message2)
// var message3 = <-messages
// fmt.Println(message3)

// case 4
// runtime.GOMAXPROCS(2)

// var messages = make(chan string)

// for _, each := range []string{"Eko", "Joko", "Budi"} {
// 	go func(who string) {
// 		var data = fmt.Sprintf("Hello %s", who)
// 		// kirim data ke channel
// 		messages <- data
// 	}(each)
// }

// for i := 0; i < 3; i++ {
// 	PrintMessage2(messages)
// }

// case 5 bufferchannnel

// runtime.GOMAXPROCS(2)

// messages := make(chan int, 3)

// go func() {
// 	for {
// 		i := <-messages
// 		fmt.Println("receive data", i)
// 	}
// }()

// for i := 0; i < 5; i++ {
// 	fmt.Println("send data", i)
// 	messages <- i
// }

// time.Sleep(2 * time.Second)

// case 6
// channel Select

// runtime.GOMAXPROCS(2)

// var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
// fmt.Println("numbers :", numbers)

// var ch1 = make(chan float64)
// go GetAverage(numbers, ch1)

// var ch2 = make(chan int)
// go GetMax(numbers, ch2)

// for i := 0; i < 2; i++ {
// 	select {
// 	case avg := <-ch1:
// 		fmt.Printf("Avg \t: %.2f \n", avg)
// 	case max := <-ch2:
// 		fmt.Printf("Max \t: %d \n", max)
// 	}
// }

// case 7
// channel close

// runtime.GOMAXPROCS(2)

// var messages = make(chan string)
// go SendMessage3(messages)
// PrintMessage3(messages)

// case 8
// channel timeout

// runtime.GOMAXPROCS(2)

// var messages = make(chan int)

// go SendDataTimeOut(messages)
// RetreiveData(messages)

// }

func RetreiveData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Print(`receive data "`, data, `"`, "\n")
		case <-time.After(time.Second * 5):
			fmt.Println("timeout. no activities under 5 seconds")
			break loop
		}
	}
}

func SendDataTimeOut(ch chan<- int) {
	// randomizer := rand.New(rand.NewSource(time.Now().Unix()))

	// for i := 0; true; i++ {
	// 	ch <- i
	// 	time.Sleep(time.Duration(randomizer.Int()%10+1) * time.Second)
	// }
}

func SendMessage3(ch chan<- string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("data %d", i)
	}
	close(ch)
}

func PrintMessage3(ch <-chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}

func GetAverage(numbers []int, ch chan float64) {
	var sum = 0
	for _, e := range numbers {
		sum += e
	}
	ch <- float64(sum) / float64(len(numbers))
}

func GetMax(numbers []int, ch chan int) {
	var max = numbers[0]
	for _, e := range numbers {
		if max < e {
			max = e
		}
	}
	ch <- max
}

func PrintMessage2(what chan string) {
	fmt.Println(<-what)
}

func Print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println(message)
		// time.Sleep(500 * time.Millisecond)
	}

}

func PrintMessage(message string) {
	for i := 0; i < 5; i++ {
		fmt.Println(message)
		time.Sleep(500 * time.Millisecond)
	}
}
