package main

import (
	"flag"
	"fmt"

	userlist "github.com/KTaclysme/tp1/user_list"
)

func main() {
	var nom string
	var action string
	var tel string

	flag.StringVar(&nom, "nom", "", "put a name")
	flag.StringVar(&action, "action", "", "an action: ajouter, modifier, supprimer, rechercher")
	flag.StringVar(&tel, "tel", "", "put a phone number")
	flag.Parse()

	fn, ok := userlist.Actions[userlist.Action(action)]
	if !ok {
		fmt.Println("Action inconnue")
		return
	}

	fn(userlist.Nom(nom), userlist.Tel(tel))
}
