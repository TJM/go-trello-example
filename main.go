package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

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
		// Pick one randomly to display
		i := rand.Intn(len(boards))
		board := boards[i]
		fmt.Printf("* %v (%v)\n", board.Name, board.ShortURL)

		// Board Lists
		lists, err := board.Lists()
		if err != nil {
			log.Fatal(err)
		}

		for _, list := range lists {
			fmt.Println("   - ", list.Name)

			// List Cards
			cards, _ := list.Cards()
			for _, card := range cards {
				fmt.Println("      + ", card.Name)
			}
		}
	}
}
