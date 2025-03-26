package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// excecutes other programs
func main() {
	fmt.Println("Welcome to the notes tool! \n")

	// Getting collection name and filename from the argument given in the command line
	collectionName := os.Args[1]
	fileName := os.Args[1] + ".txt"

	if _, err := os.Stat(fileName); err == nil {
		fmt.Println("You have selected collection " + collectionName)

		// If the collection does not exist, create one
	} else if os.IsNotExist(err) {
		command := exec.Command("touch", fileName)
		err := command.Run() // Execute the command
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		fmt.Println("New collection " + collectionName + " created")
	}
	start(fileName)

}

// loadNotes loads notes from a specified file
func LoadNotes(fileName string) []string {
	file, err := os.Open(fileName) // Open the file for reading
	if err != nil {
		if os.IsNotExist(err) {
			return []string{} // If file does not exist, return an empty slice
		}
		fmt.Printf("Error opening file: %v\n", err)
		return []string{}
	}
	defer file.Close()

	var notes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // Read each line of the file into the notes slice
		notes = append(notes, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	return notes
}

const (
	Reset = "\033[0m"
	Red   = "\033[31m"
)

type ActionType string

const (
	Show   ActionType = "Show"
	Add    ActionType = "Add"
	Delete ActionType = "Delete"
	Exit   ActionType = "Exit"
)

func start(fileName string) {
	for {
		notes := LoadNotes(fileName)
		action, message := GetInput()

		switch action {
		case Show:
			showNotes(notes)
		case Add:
			addNote(notes, fileName)
		case Delete:
			deleteNote(notes, fileName)
		case Exit:
			fmt.Println("Thank you for using the notes tool! Goodbye.")
			return
		default:
			fmt.Println("Unknown action:", message)
		}
	}
}

func GetInput() (ActionType, string) {
	getAction := GetNumber(
		"\nSelect action (1-4):"+
			"\n1. Show"+
			"\n2. Add"+
			"\n3. Delete"+
			"\n4. Exit"+
			"\n", 1, 4)

	switch getAction {
	case 1:
		return Show, "Show"
	case 2:
		return Add, "Add"
	case 3:
		return Delete, "Delete"
	case 4:
		return Exit, "Exit"
	default:
		return "", "Invalid"
	}
}

func GetNumber(question string, lowerLimit int, maxLimit int) int {
	for {
		text := GetText(question) //false)
		number, err := strconv.Atoi(text)
		if err != nil {
			fmt.Printf("\n%sError: Input: Only numbers are allowed!%s\n", Red, Reset)
			fmt.Println("Please enter a valid number or 0 to cancel.")
			continue
		}

		if number <= maxLimit && number >= lowerLimit {
			return number
		} else {
			fmt.Printf("\n%sError: There is no such option!%s\n", Red, Reset)
			fmt.Println("Please select an operation from the menu or exit the program.")
		}

	}
}

func GetText(question string) string {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(question)
		scanner.Scan()

		inputText := scanner.Text()
		inputText = strings.Trim(inputText, " ")

		if inputText != "" {
			return inputText
		}

		fmt.Printf("%sError: Empty input!%s", Red, Reset)
		fmt.Println("\nPlease enter your note.")
	}
}

//addNote, add notes

func addNote(notes []string, fileName string) {
	note := GetText("Enter your note: ") //gets the note from user input
	//notes = append(notes, note)                        //adds a note to the existing note
	bash_command := "echo " + note + " >> " + fileName // writes a note to the text file
	command := exec.Command("/bin/sh", "-c", bash_command)
	err := command.Run() // Execute the command
	if err != nil {
		fmt.Println("Error adding to notes:", err)
	}
	return
}

// show notes
func showNotes(notes []string) { //cheks if the length of notes is 0, slice
	if len(notes) == 0 {
		fmt.Println("There are no notes in your collection yet.")
	} else {
		fmt.Println("Notes:") //Print notes + 1 (next note)
		for i, note := range notes {
			fmt.Printf("%1d - %s\n", i+1, note) // %1d (number of the note), %s text of the note
		}
	}
}

// loadNotes loads notes from a specified file
func loadNotes(fileName string) []string {
	file, err := os.Open(fileName) // Open the file for reading
	if err != nil {
		if os.IsNotExist(err) {
			return []string{} // If file does not exist, return an empty slice
		}
		fmt.Printf("Error opening file: %v\n", err)
		return []string{}
	}
	defer file.Close()

	var notes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // Read each line of the file into the notes slice
		notes = append(notes, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	return notes
}

//delete notes

func deleteNote(notes []string, fileName string) {
	if len(notes) == 0 { //Are there notes to delete?
		fmt.Println("No available notes.")
		return
	}

	//display available notes
	fmt.Println("Enter the number of note to delete. 0 - cancel:")
	showNotes(notes)

	index := GetNumber("", 0, len(notes)) //userinput to delete

	if index == 0 {
		fmt.Println("Deletion is cancelled")
	}

	index-- //convert to zero based index

	notes = append(notes[:index], notes[index+1:]...) //remove notes from slice

	// replace the notes in the file
	bash_command := "> " + fileName // deletes the contents of the file
	command := exec.Command("/bin/sh", "-c", bash_command)
	err := command.Run() // Execute the command
	if err != nil {
		fmt.Println("Error deleting a note:", err)
	}
	for _, note := range notes {
		bash_command = "echo " + note + " >> " + fileName // writes a note to the text file
		command := exec.Command("/bin/sh", "-c", bash_command)
		err := command.Run() // Execute the command
		if err != nil {
			fmt.Println("Error deleting a note:", err)
		}
	}

	fmt.Println("Note deleted.")
	return

}
