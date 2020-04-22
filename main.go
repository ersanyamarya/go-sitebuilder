package main

// Importing required packages
import (
	"log"
)

func main() {
	var pages Page
	err := pages.GetChaptersFromFile("mqtt.json")
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range pages.Content {
		p.CreateHTML()
	}
}
