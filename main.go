package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/TJM/go-trello"
	"github.com/alexflint/go-arg"
)

func main() {
	// Parse Command Line Args
	var args struct {
		AppKey string `arg:"required,env:TRELLO_APP_KEY" help:"Trello API App Key, Obtain yours at https://trello.com/app-key ... (env: TRELLO_APP_KEY)"`
		Token  string `arg:"required,env:TRELLO_TOKEN" help:"Trello API App Key, Authorize your App Key to use your account at <https://trello.com/1/connect?key=<appKey from above>&name=Go-Trello-Example-delete-boards&response_type=token&scope=read,write&expiration=1day> (env: TRELLO_TOKEN)"`
	}
	arg.MustParse(&args)

	// New Trello Client
	appKey := args.AppKey
	token := args.Token
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
