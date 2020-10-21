# Golang Trello Example API client

go-trello-example is a [Go](http://golang.org/) example client package for accessing the [Trello](http://www.trello.com/) [API](https://developer.atlassian.com/cloud/trello/guides/rest-api/api-introduction/) using the [TJM/go-trello](https://github.com/TJM/go-trello) module.

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
./go-trello-example

```

## License

Licensed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0).
