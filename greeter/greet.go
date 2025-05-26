package greeter

import "fmt"

type Language string

var phrasebook = map[Language]string{
	"en": "Hello world",
	"fr": "Bonjour le monde",
	"de": "Hallo welt",
	"":   "Unsupported language",
	"es": "Unsupported language",
}

func Greet(l Language) (string, error) {
	greeting, ok := phrasebook[l]
	if !ok {
		return "", fmt.Errorf("unsupported language %q", l)
	}
	return greeting, nil
}
