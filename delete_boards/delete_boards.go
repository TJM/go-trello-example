package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/TJM/go-trello"
	"github.com/alexflint/go-arg"
)

var args struct {
	AppKey     string `arg:"required,env:TRELLO_APP_KEY" help:"Trello API App Key.\n\t\t Obtain yours at https://trello.com/app-key\n\t\t (env: TRELLO_APP_KEY)"`
	Token      string `arg:"required,env:TRELLO_TOKEN" help:"Trello API App Key.\n\t\t Authorize your App Key to use your account at <https://trello.com/1/connect?key=<appKey from above>&name=Go-Trello-Example-delete_boards&response_type=token&scope=read,write&expiration=1day>\n\t\t (env: TRELLO_TOKEN)"`
	AnyOf      bool   `help:"Match AnyOf the StartsWith, Contains or EndsWith conditions. By default board name must match all of the conditions."`
	StartsWith string `help:"Select boards to delete that *start with* this string"`
	Contains   string `help:"Select boards to delete that *contain* this string"`
	EndsWith   string `help:"Select boards to delete that *end with* this string"`
	Delete     bool   `help:"Actually DELETE the boards (defaults to false so you can see what will happen)"`
	Debug      bool   `help:"Enable debugging output"`
}

var w *tabwriter.Writer

func main() {
	// Parse Command Line Args
	arg.MustParse(&args)

	// Tab Writer
	w = new(tabwriter.Writer)
	w.Init(os.Stdout, 4, 4, 2, ' ', tabwriter.TabIndent)

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
	fmt.Printf("Trello User: %s (%s) <%s>\n", user.FullName, user.Username, user.URL)

	// @trello Boards
	boards, err := user.Boards()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Trello Boards: %v\n", len(user.IDBoards))
	fmt.Fprintf(w, "\tAction\tBoard Name\tURL\n")
	fmt.Fprintf(w, "\t------\t----------\t---\n")

	for _, board := range boards {
		var action string
		if args.Debug {
			boardJSON, err := json.MarshalIndent(board, "", "  ")
			if err != nil {
				log.Fatalf(err.Error())
			}
			fmt.Printf("====================\n%s:\n%s\n\n", board.Name, string(boardJSON))
		}
		if args.AnyOf {
			// AnyOf
			if args.StartsWith != "" && strings.HasPrefix(board.Name, args.StartsWith) {
				action = removeBoard(board, user)
			} else if args.Contains != "" && strings.Contains(board.Name, args.Contains) {
				action = removeBoard(board, user)
			} else if args.EndsWith != "" && strings.HasSuffix(board.Name, args.EndsWith) {
				action = removeBoard(board, user)
			} else {
				action = "KEEP"
			}
		} else {
			if args.StartsWith == "" && args.Contains == "" && args.EndsWith == "" { // If no conditions are set, KEEP
				action = "KEEP"
			} else if (args.StartsWith == "" || strings.HasPrefix(board.Name, args.StartsWith)) &&
				(args.Contains == "" || strings.Contains(board.Name, args.Contains)) &&
				(args.EndsWith == "" || strings.HasSuffix(board.Name, args.EndsWith)) { // If a condition is set, and matches, DELETE
				action = removeBoard(board, user)
			} else { // KEEP by default
				action = "KEEP"
			}
		}
		fmt.Fprintf(w, "\t%s\t%s\t<%s>\n", action, board.Name, board.ShortURL)
	}

	// Output Tabwriter Table
	fmt.Printf("\n")
	w.Flush()

	if !args.Delete {
		fmt.Printf("\n\n ** Run again with --delete flag to actually delete the board(s).\n\n")
	}
}

func removeBoard(board *trello.Board, user *trello.Member) (action string) {
	if args.Delete {
		if board.IsAdmin(user) {
			fmt.Printf("DELETING: %s <%s> ...\n", board.Name, board.ShortURL)
			err := board.Delete()
			action = "DELETED"
			if err != nil {
				fmt.Printf("ERROR Deleting board: %v\n", err)
				action = "ERROR DELETING"
			}
		} else {
			fmt.Printf("LEAVING:  %s <%s> ...\n", board.Name, board.ShortURL)
			err := board.RemoveMember(user)
			action = "LEFT"
			if err != nil {
				fmt.Printf("ERROR Leaving board: %v\n", err)
				action = "ERROR LEAVING"
			}
		}
	} else {
		if board.IsAdmin(user) {
			action = "DELETE"
		} else {
			action = "LEAVE"
		}
	}
	return
}
