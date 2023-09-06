package main

import (
	"golang.org/x/text/message"
)

func main() {
	// Run code affected by vulns in golang.org/x/text.
	p := message.NewPrinter(message.MatchLanguage("en"))
	p.Println(123456.78)
}
