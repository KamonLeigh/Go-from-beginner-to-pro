package main

import (
	"fmt"
	"os"
)

func getUsers() map[string]string {
	users := map[string]string{
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
		"073": "Tracy",
	}

	return users
}

func getUser(id string) (string, bool) {

	users := getUsers()

	name, exists := users[id]

	return name, exists
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("User ID not passed")
		os.Exit(1)
	}
	userID := os.Args[1]
	name, exists := getUser(userID)

	if !exists {
		fmt.Printf("Passed user ID (%v) not found. \nUsers:\n", userID)

		for id, name := range getUsers() {
			fmt.Println(" ID:", id, "Name:", name)
		}

		os.Exit(1)
	}

	fmt.Println("Hi", name)

}
