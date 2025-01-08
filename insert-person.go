package main

// type Person struct {
// 	Id          int
// 	Name        string
// 	Address     string
// 	PhoneNumber string
// 	Email       string
// }

// func clearConsole() {
// 	fmt.Print("\033[H\033[2J")
// }

// func main() {

// 	// connection database
// 	dsn := "host=localhost user=postgres password=1234 dbname=person_db port=5432 sslmode=disable TimeZone=Asia/Jakarta"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Failed to connect to database: %v", err)
// 	}

// 	// connet to database response

// 	fmt.Println("Connected to database!")

// 	if err := db.AutoMigrate(&Person{}); err != nil {
// 		log.Fatalf("Failed to migrate database: %v", err)
// 	}

// 	fmt.Println("Database migrated !")

// 	fmt.Println("===============================================================")
// 	fmt.Println("===============================================================")
// 	fmt.Println("===============================================================")
// 	fmt.Println("================= WELCOME TO APP INPUT PERSON =================")
// 	fmt.Println("===============================================================")
// 	fmt.Println("===============================================================")
// 	fmt.Println("===============================================================")

// 	var p1 Person
// 	var personBulk []Person

// 	var name, address, phoneNumber, email, continueInput string

// 	for {

// 		// clearConsole()
// 		// Input data
// 		fmt.Print("Name: ")
// 		fmt.Scanln(&name)
// 		fmt.Print("Address: ")
// 		fmt.Scanln(&address)
// 		fmt.Print("Phone Number: ")
// 		fmt.Scanln(&phoneNumber)
// 		fmt.Print("Email: ")
// 		fmt.Scanln(&email)

// 		// Assign data to struct
// 		p1 = Person{
// 			Name:        name,
// 			Address:     address,
// 			PhoneNumber: phoneNumber,
// 			Email:       email,
// 		}

// 		// Append to array
// 		personBulk = append(personBulk, p1)

// 		// Ask user to continue
// 		fmt.Print("Continue? (y/n): ")
// 		fmt.Scanln(&continueInput)
// 		if continueInput == "n" || continueInput == "N" {
// 			break
// 		}
// 	}

// 	if len(personBulk) > 0 {
// 		// Save data to database

// 		if err := db.Create(&personBulk).Error; err != nil {
// 			log.Fatalf("Failed to save data to database: %v", err)
// 		} else {
// 			fmt.Println("Data saved to database successfully!")
// 		}
// 	} else {
// 		fmt.Println("No data to save.")
// 	}

// 	// Print all collected data
// 	// clearConsole()
// 	// fmt.Println("\nData Collected from Database:")
// 	// var persons []Person
// 	// if err := db.Find(&persons).Error; err != nil {
// 	// 	log.Printf("Failed to retrieve data: %v", err)
// 	// } else {
// 	// 	for i, person := range persons {
// 	// 		fmt.Printf("Person %d: %+v\n", i+1, person)
// 	// 	}
// 	// }

// }
