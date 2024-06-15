package main

import (
	"log"

	"github.com/rosencsd/planet_defense/pkg/cmd"
)

func main() {
	if err := cmd.NewHammurabiCmd().Execute(); err != nil {
		log.Println(err)
	}
}
