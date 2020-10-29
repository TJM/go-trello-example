package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/TJM/go-trello"
)

func main() {
	// New Trello Client
	appKey := os.Getenv("TRELLO_APP_KEY")
	token := os.Getenv("TRELLO_TOKEN")
	trello, err := trello.NewAuthClient(appKey, &token)
	if err != nil {
		log.Fatal(err)
	}

	// User @trello
	user, err := trello.Member("me")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Trello User: %v\n", user.FullName)

	// @trello Boards
	boards, err := user.Boards()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Trello Boards: %v\n", len(user.IDBoards))

	if len(boards) > 0 {
		// Delete any board with "GoTrelloTest" in the name
		for _, board := range boards {
			if strings.Contains(board.Name, "GoTestTrello") {
				fmt.Println("Delete: " + board.Name)
				//board.Delete() // Uncomment this to make it actually delete
			} else {
				fmt.Println("Keep: " + board.Name)
			}

		}
	}
}
