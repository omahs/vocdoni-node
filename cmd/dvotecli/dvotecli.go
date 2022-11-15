package main

import (
	"go.vocdoni.io/dvote/cmd/dvotecli/commands"
)

func main() {
	commands.Execute()
}

func deadFunc() {
	// no-op, to trigger staticcheck failure
}
