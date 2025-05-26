package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/KTaclysme/helloworld/greeter"
)

func main() {
	var lang string
	flag.StringVar(&lang, "lang", "en", "The required language: en, fr, de, ...")
	flag.Parse()
	greeting, err := greeter.Greet(greeter.Language(lang))
	if err != nil {
		log.Fatalf("Error greeting: %v", err)
	}
	fmt.Println(greeting)
}
