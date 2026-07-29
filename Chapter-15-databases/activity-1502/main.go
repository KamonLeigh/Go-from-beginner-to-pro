package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

type Messages struct {
	userID  int
	message string
}

func main() {
	var toLookFor string
	var message string
	var email string
	var name string

	messages := []Messages{
		{1, "Hi Florian, when are you coming home?"},
		{1, "Can you send some cash?"},
		{2, "Hi can you bring some bread and milk?"},
		{7, "Well..."},
	}

	db, err := sql.Open("postgres", "user=postgres password=Start!123 host=127.0.0.1 port=5432 dbname=postgresgo sslmode=disable")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("The connetion to the DB was successfully initialised")
	}

	TableCreate := `
	CREATE TABLE IF NOT EXISTS  messages
	(
    UserID integer NOT NULL,
    Message character varying(280) COLLATE pg_catalog."default" NOT NULL
	)
	WITH (
    OIDS = FALSE
	)
	TABLESPACE pg_default;

	ALTER TABLE public.messages
    OWNER to postgres;
	`

	_, err = db.Exec(TableCreate)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("The table called  was successfully created")
	}

	insert, insertErr := db.Prepare("INSERT INTO messages (UserID, Message) VALUES($1, $2)")

	if insertErr != nil {
		panic(insertErr)
	}

	for _, message := range messages {

		_, err := insert.Exec(message.userID, message.message)

		if err != nil {
			panic(err)
		} else {
			fmt.Println("The UserID: ", message.userID, " with message: ", message.message)

		}
	}

	insert.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Give me the user's name: ")

	toLookFor, err = reader.ReadString('\n')
	toLookFor = strings.Trim(toLookFor, "\r\n")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Looking for all the messages of user with name:", toLookFor, "##")
	}

	UserMessages := "SELECT users.Name, users.Email, messages.Message FROM messages INNER JOIN users ON users.ID=messages.UserID WHERE users.Name LIKE $1"

	userMessages, err := db.Prepare(UserMessages)

	if err != nil {
		panic(err)
	}

	result, err := userMessages.Query(toLookFor)

	if err != nil {
		panic(err)
	}

	numberof := 0

	for result.Next() {
		numberof++
	}

	if numberof == 0 {
		fmt.Println("The query returned nothing, no such user:", toLookFor)
	} else {
		fmt.Println("There are a total of", numberof, "messages from the user:", toLookFor)

		result, err = userMessages.Query(toLookFor)

		if err != nil {
			panic(err)
		}

		for result.Next() {
			err = result.Scan(&message, &email, &name)

			if err != nil {
				panic(err)
			}

			fmt.Println("The user:", name, "with email:", email, "has sent the following message:", message)

		}
	}

	userMessages.Close()
	db.Close()

}
