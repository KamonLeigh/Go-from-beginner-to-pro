package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

var (
	ErrorWorkingFileNotFound = errors.New("The working file is not Found.")
)

func createBackup(working, backup string) error {
	_, err := os.Stat(working)

	if err != nil {
		if os.IsNotExist(err) {
			return ErrorWorkingFileNotFound
		}

		return err
	}

	workFile, err := os.Open(working)

	if err != nil {
		return err
	}

	content, err := io.ReadAll(workFile)

	if err != nil {
		return err
	}

	err = os.WriteFile(backup, content, 0644)

	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func addNote(workingFile, notes string) error {
	notes += "\n"

	f, err := os.OpenFile(workingFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	defer f.Close()

	if _, err := f.Write([]byte(notes)); err != nil {
		return err
	}

	return nil
}

func main() {
	backupFile := "backupFile.txt"
	workingFile := "note.txt"
	data := "note"

	err := createBackup(workingFile, backupFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 0; i <= 10; i++ {
		note := data + " " + strconv.Itoa(i)

		err := addNote(workingFile, note)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
