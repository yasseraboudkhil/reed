package main

import (
	"github.com/yasseraboudkhil/reed/cmd"
	"log"
)

func main() {

	if err := cmd.RootCommand().Execute(); err != nil {
		log.Fatal("Could not get a root command")
	}
}
