package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	name  string
	email string
}

type Message struct {
	useerID string
	message string
}

func main() {
	users := []User{
		User{
			name:  "Szabo Daniel",
			email: "daniel@pakt.com",
		},
		User{
			name:  "Szabo Florian",
			email: "florian@pakt.com",
		},
	}

	db, err := sql.Open("postgres", "user=postgres password=Start!123 host=127.0.0.1 port=5432 dbname=postgresgo sslmode=disable")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("The connetion to the DB was successfully initialised")
	}

	TableCreate := `
		CREATE TABLE IF NOT EXISTS Users
		(
			Id SERIAL PRIMARY KEY,
			Name varchar(100) NOT NULL UNIQUE,
			Email varchar(100) NOT NULL
		)
		WITH (
		OIDS = FALSE
		)
		TABLESPACE pg_default;

		ALTER TABLE Users
    	OWNER to postgres;

	`
	_, err = db.Exec(TableCreate)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("The table called Users was successfully created")
	}

	insert, insertErr := db.Prepare("INSERT INTO users (Name, Email) VALUES($1, $2)")

	if insertErr != nil {
		panic(insertErr)
	}

	for _, user := range users {

		_, err := insert.Exec(user.name, user.email)

		if err != nil {
			panic(err)
		} else {
			fmt.Println("The user with name: ", user.name, " and email: ", user.email, " was successfully added!")
		}

	}

	insert.Close()

	update := "UPDATE users SET Email=$1 WHERE id=$2"

	_, err = db.Exec(update, "daniel1@pakt.com", 1)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("The user's email address was succefully updated!")
	}

	remove := "DELETE FROM users WHERE name=$1"
	_, err = db.Exec(remove, users[1].name)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("The second user was succeesfully removed!")
	}

	db.Close()

}
