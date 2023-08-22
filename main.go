package main

import (
	"log"

	"github.com/nicolaslh/ntts/commands"
)

func main() {

	log.SetFlags(0)
	err := commands.Execute()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}
