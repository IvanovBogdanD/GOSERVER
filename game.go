package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Room struct {
	name  string
	look  string
	items map[string]bool
	exits []string
}

var rooms map[string]*Room
var current *Room
var inventory map[string]bool
var backpack bool
var doorOpen bool

func initGame() {
	rooms = map[string]*Room{
		"кухня": {
			name:  "кухня",
			look:  "ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор",
			items: map[string]bool{},
			exits: []string{"коридор"},
		},
		"коридор": {
			name:  "коридор",
			look:  "ничего интересного. можно пройти - кухня, комната, улица",
			items: map[string]bool{},
			exits: []string{"кухня", "комната", "улица"},
		},
		"комната": {
			name:  "комната",
			look:  "ты в своей комнате. можно пройти - коридор",
			items: map[string]bool{"ключи": true, "конспекты": true, "рюкзак": true},
			exits: []string{"коридор"},
		},
		"улица": {
			name:  "улица",
			look:  "на улице весна. можно пройти - домой",
			items: map[string]bool{},
			exits: []string{"домой"},
		},
	}
	current = rooms["кухня"]
	inventory = map[string]bool{}
	backpack = false
	doorOpen = false
}

func handleCommand(cmd string) string {
	parts := strings.Split(cmd, " ")
	switch parts[0] {
	case "осмотреться":
		return lookAround()
	case "идти":
		if len(parts) < 2 {
			return "нет пути"
		}
		return goTo(parts[1])
	case "надеть":
		if len(parts) < 2 {
			return "нет такого"
		}
		return wear(parts[1])
	case "взять":
		if len(parts) < 2 {
			return "нет такого"
		}
		return take(parts[1])
	case "применить":
		if len(parts) < 3 {
			return "не к чему применить"
		}
		return use(parts[1], parts[2])
	default:
		return "неизвестная команда"
	}
}

func lookAround() string {
	if current.name == "комната" {
		if len(current.items) == 0 {
			return "пустая комната. можно пройти - коридор"
		}
		s := []string{}
		if current.items["ключи"] || current.items["конспекты"] {
			t := []string{}
			if current.items["ключи"] {
				t = append(t, "ключи")
			}
			if current.items["конспекты"] {
				t = append(t, "конспекты")
			}
			s = append(s, "на столе: "+strings.Join(t, ", "))
		}
		if current.items["рюкзак"] {
			s = append(s, "на стуле: рюкзак")
		}
		return strings.Join(s, ", ") + ". можно пройти - коридор"
	}
	if current.name == "кухня" {
		if backpack {
			return "ты находишься на кухне, на столе: чай, надо идти в универ. можно пройти - коридор"
		}
		return rooms["кухня"].look
	}
	return current.look
}

func goTo(place string) string {
	for _, e := range current.exits {
		if e == place {
			if place == "улица" && !doorOpen {
				return "дверь закрыта"
			}
			current = rooms[place]
			if current.name == "кухня" && backpack {
				return "кухня, ничего интересного. можно пройти - коридор"
			}
			if current.name == "комната" {
				return "ты в своей комнате. можно пройти - коридор"
			}
			return lookAround()
		}
	}
	return "нет пути в " + place
}

func wear(item string) string {
	if item == "рюкзак" && current.items["рюкзак"] {
		backpack = true
		delete(current.items, "рюкзак")
		return "вы надели: рюкзак"
	}
	return "нет такого"
}

func take(item string) string {
	if !backpack {
		return "некуда класть"
	}
	if !current.items[item] {
		return "нет такого"
	}
	delete(current.items, item)
	inventory[item] = true
	return "предмет добавлен в инвентарь: " + item
}

func use(item, target string) string {
	if !inventory[item] {
		return "нет предмета в инвентаре - " + item
	}
	if item == "ключи" && target == "дверь" {
		doorOpen = true
		return "дверь открыта"
	}
	return "не к чему применить"
}

func main() {
	initGame()
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Добро пожаловать! Пиши команды (\"выход\" чтобы выйти)")
	for {
		fmt.Print("> ")
		if !reader.Scan() {
			break
		}
		cmd := reader.Text()
		if cmd == "выход" {
			fmt.Println("Пока!")
			break
		}
		fmt.Println(handleCommand(cmd))
	}
}
