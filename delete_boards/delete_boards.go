package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/TJM/go-trello"
	"github.com/alexflint/go-arg"
)

var args struct {
	AppKey     string `arg:"required,env:TRELLO_APP_KEY" help:"Trello API App Key, Obtain yours at https://trello.com/app-key ... (env: TRELLO_APP_KEY)"`
	Token      string `arg:"required,env:TRELLO_TOKEN" help:"Trello API App Key, Authorize your App Key to use your account at <https://trello.com/1/connect?key=<appKey from above>&name=Go-Trello-Example-delete-boards&response_type=token&scope=read,write&expiration=1day> (env: TRELLO_TOKEN)"`
	AnyOf      bool   `help:"Match AnyOf the StartsWith, Contains or EndsWith conditions. By default board name must match all of the conditions."`
	StartsWith string `help:"Select boards to delete that *start with* this string"`
	Contains   string `help:"Select boards to delete that *contain* this string"`
	EndsWith   string `help:"Select boards to delete that *end with* this string"`
	Delete     bool   `help:"Actually DELETE the boards (defaults to false so you can see what will happen)"`
}

func main() {
	// Parse Command Line Args

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

	for _, board := range boards {
		if args.AnyOf {
			// AnyOf
			if args.StartsWith != "" && strings.HasPrefix(board.Name, args.StartsWith) {
				fmt.Println("Delete: " + board.Name)
				removeBoard(board)
			} else if args.Contains != "" && strings.Contains(board.Name, args.Contains) {
				fmt.Println("Delete: " + board.Name)
				removeBoard(board)
			} else if args.EndsWith != "" && strings.HasSuffix(board.Name, args.EndsWith) {
				fmt.Println("Delete: " + board.Name)
				removeBoard(board)
			} else {
				fmt.Println("Keep: " + board.Name)
			}
		} else {
			if args.StartsWith == "" && args.Contains == "" && args.EndsWith == "" { // If no conditions are set, KEEP
				fmt.Println("Keep: " + board.Name)
			} else if (args.StartsWith == "" || strings.HasPrefix(board.Name, args.StartsWith)) &&
				(args.Contains == "" || strings.Contains(board.Name, args.Contains)) &&
				(args.EndsWith == "" || strings.HasSuffix(board.Name, args.EndsWith)) { // If a condition is set, and matches, DELETE
				fmt.Println("Delete: " + board.Name)
				removeBoard(board)
			} else { // KEEP by default
				fmt.Println("Keep: " + board.Name)
			}
		}
	}

	if !args.Delete {
		fmt.Printf("\n\n ** Run again with --delete flag to actually delete the boards.\n\n")
	}
}

func removeBoard(board trello.Board) {
	if args.Delete {
		err := board.Delete()
		if err != nil {
			fmt.Printf("ERROR Deleting board: %v\n", err)
		}
	}
}
