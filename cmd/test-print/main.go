package main

import (
	"golang.org/x/text/message"

	types "github.com/adevinta/vulcan-types"
)

func main() {
	// Run code affected by vulns in golang.org/x/text.
	p := message.NewPrinter(message.MatchLanguage("en"))
	p.Println(123456.78)

	types.IsDockerImage("go1.21")
}
