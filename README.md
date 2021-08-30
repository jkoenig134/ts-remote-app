# ts5 Remote Application Library

This library is wrapping the teamspeak remote application API.

The library is in a really early state and not even close completed ... contributions are welcome!

Most information about the interface is from [this](https://community.teamspeak.com/t/teamspeak-5-0-0-beta54-x/22988)
and [this](https://community.teamspeak.com/t/teamspeak-5-0-0-beta55/23318) official blog post. The other information
comes from trial and error and reading events that are executed. Contact me if you encounter events that are not handled
by my lib.

## Installation

simply run `go get github.com/jkoenig134/ts-remote-app`

## Usage

`example/main.go` provides a good example on how to listen to events and communicate to the client.

