# Golang Trello Example - Delete Boards

go-trello-example/delete_boards is a [Go](http://golang.org/) *useful* example client package for accessing the [Trello](http://www.trello.com/) [API](https://developer.atlassian.com/cloud/trello/guides/rest-api/api-introduction/) using the [TJM/go-trello](https://github.com/TJM/go-trello) module to delete trello boards that match specific conditions

## Example

Prerequisites:

* Retrieve your `appKey`: <https://trello.com/app-key> (NOTE: This identifies "you" as the developer of the application)
* Retrieve your (temporary) `token`: <https://trello.com/1/connect?key=MYKEYFROMABOVE&name=MYAPPNAME&response_type=token&scope=read,write&expiration=1day>

Building:

```bash
go build
```

Running: (using your appKey and token from above)

```bash
export TRELLO_APP_KEY=xxxxxxxxxxxxxxxxxxx
export TRELLO_TOKEN=yyyyyyyyyyyyyyyyyyyy
./delete_boards --help
Usage: delete_boards --appkey APPKEY --token TOKEN [--anyof] [--startswith STARTSWITH] [--contains CONTAINS] [--endswith ENDSWITH] [--delete] [--debug]

Options:
  --appkey APPKEY        Trello API App Key.
         Obtain yours at https://trello.com/app-key
         (env: TRELLO_APP_KEY)
  --token TOKEN          Trello API App Key.
         Authorize your App Key to use your account at <https://trello.com/1/connect?key=<appKey from above>&name=Go-Trello-Example-delete_boards&response_type=token&scope=read,write&expiration=1day>
         (env: TRELLO_TOKEN)
  --anyof                Match AnyOf the StartsWith, Contains or EndsWith conditions. By default board name must match all of the conditions.
  --startswith STARTSWITH
                         Select boards to delete that *start with* this string
  --contains CONTAINS    Select boards to delete that *contain* this string
  --endswith ENDSWITH    Select boards to delete that *end with* this string
  --delete               Actually DELETE the boards (defaults to false so you can see what will happen)
  --debug                Enable debugging output
  --help, -h             display this help and exit

```

NOTE: The `--delete` flag is required to *actually* delete boards, so you can safely test this.

```bash
[tmcneely@local delete_boards] $ ./delete_boards --contains Patching
Trello User: Tommy McNeely
Trello Boards: 20
Keep:   AFTER: CWOW-DEV System Updates 2020Q3
Keep:   CWOW-DEV System Updates 2020Q3
Keep:   Morpheus Features
Delete: Patching 2019-09-03
Delete: Patching 2019-10-01
Delete: Patching 2019-10-29
Delete: Patching 2019-11-12
Delete: Patching 2019-12-08
Delete: Patching 2020-1-22
Delete: Patching 2020-1-9
Delete: Patching 2020-2-5
Delete: Patching 2020-4-15
Delete: Patching 2020-4-29
Delete: Patching 2020-5-20
Delete: Patching 2020-7-22
Delete: Patching 2020-8-30
Delete: Patching 2020-8-5
Delete: Patching 2020-9-30
Keep:   SDLC Prod 10/15
Keep:   Welcome Board
```

## License

Licensed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0).
