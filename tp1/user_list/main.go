package userlist

import (
	"encoding/json"
	"fmt"
	"os"
)

type Action string
type Nom string
type Tel string

var Actions = map[Action]func(nom Nom, tel Tel){
	"ajouter":    ajouter,
	"rechercher": rechercher,
	"modifier":   modifier,
	"supprimer":  supprimer,
}
var users = make(map[Nom]Tel)

func load() {
	data, err := os.ReadFile("users.json")
	if err == nil {
		json.Unmarshal(data, &users)
	}
}

func save() {
	data, _ := json.MarshalIndent(users, "", "  ")
	os.WriteFile("users.json", data, 0644)
}

func ajouter(nom Nom, tel Tel) {
	load()
	if _, exists := users[nom]; exists {
		fmt.Printf("User %s already exists with phone number %s\n", nom, users[nom])
	} else {
		users[nom] = tel
		save()
		fmt.Printf("User %s added with phone number %s\n", nom, tel)
	}
}

func rechercher(nom Nom, _ Tel) {
	load()
	if tel, exists := users[nom]; exists {
		fmt.Printf("User %s found with phone number %s\n", nom, tel)
	} else {
		fmt.Printf("User %s not found\n", nom)
	}
}

func modifier(nom Nom, tel Tel) {
	load()
	if _, exists := users[nom]; exists {
		users[nom] = tel
		save()
		fmt.Printf("User %s updated with phone number %s\n", nom, tel)
	} else {
		fmt.Printf("User %s not found\n", nom)
	}
}

func supprimer(nom Nom, _ Tel) {
	load()
	if _, exists := users[nom]; exists {
		delete(users, nom)
		save()
		fmt.Printf("User %s deleted\n", nom)
	} else {
		fmt.Printf("User %s not found\n", nom)
	}
}
