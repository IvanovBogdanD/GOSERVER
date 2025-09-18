package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Gamestatus struct {
	State     string   //местоположение перса
	Inventory []string //slice of strings инвентаря, как ни странно
	Rukzak    bool     //есть - нет рюкзака
}
type Wall struct {
	Room  string   //название комнаты
	Walls []string //куда нельзя из неё пройти, читай стены
}
type Room struct {
	RoomName        string
	RoomLoot        []string
	RoomObjects     []string
	RoomDescription string
	RoomWalls       Wall
}

var game = Gamestatus{} //главная переменная игрока

var locations = []string{"кухня", "коридор", "комната", "улица"}                            //локации
var items = []string{"телефон", "конспекты", "ключи"}                                       //предметы
var actions = []string{"идти", "осмотреться", "взять", "применить", "завтракать", "надеть"} //действия
var doorlock bool

/*
эта функция инициализирует игровой мир - все комнаты
если что-то было - оно корректно перезатирается
*/
var gameobstacles = []Wall{ //препятствия
	{"кухня", []string{"кухня", "комната", "улица"}},
	{"коридор", []string{"коридор"}},
	{"комната", []string{"комната", "улица", "кухня"}},
	{"улица", []string{"комната", "улица", "кухня"}},
}
var rooms map[string]Room

func initGame() {
	rooms = map[string]Room{
		"кухня": {"кухня",
			[]string{},
			[]string{"стол"},
			"ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор",
			gameobstacles[0]},

		"коридор": {"коридор",
			[]string{},
			[]string{},
			"ничего интересного. можно пройти - кухня, комната, улица",
			gameobstacles[1]},

		"комната": {"комната",
			[]string{"ключи", "конспекты", "рюкзак"},
			[]string{"стул", "стол"},
			"ты в своей комнате.",
			gameobstacles[2]},

		"улица": {"улица",
			[]string{},
			[]string{},
			"на улице весна. можно пройти - домой",
			gameobstacles[3]},
	}

	doorlock = true
	game = Gamestatus{"кухня", //местоположение перса
		[]string{}, //slice of strings инвентаря, как ни странно
		false,      //есть - нет рюкзака
	}
}
func ObstacleCheck(room string) bool {
	number := 0
	for i := range gameobstacles {
		if gameobstacles[i].Room == room {
			number = i
			break
		}
	}
	for j := range gameobstacles[number].Walls { //проверка на стены (проходка по ним)
		if game.State == gameobstacles[number].Walls[j] {
			return false
		}
	}
	return true
}

func GameGo(command string) string {
	for i := range locations {
		if command == locations[i] && command != "улица" {
			if ObstacleCheck(command) == true {
				game.State = command
				return rooms[command].RoomDescription
			} else {
				return "нет пути " + command
			}
		} else if command == "улица" {
			if ObstacleCheck(command) == true && doorlock == false {
				game.State = command
				return rooms[command].RoomDescription
			} else if ObstacleCheck(command) == true && doorlock == true {
				return "дверь закрыта"
			} else {
				return "нет пути " + command
			}
		}
	}
	return "нет пути " + command
}

func GameLook() string {
	room := rooms[game.State]
	parts := []string{room.RoomDescription}

	if len(room.RoomLoot) > 0 {
		parts = append(parts, "на столе: "+strings.Join(room.RoomLoot, ", "))
	}

	if game.Rukzak && contains(room.RoomLoot, "рюкзак") {
		parts = append(parts, "вы надели: рюкзак")
	}

	return strings.Join(parts, ". ")
}
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func GameUse(command string) string {
	if game.State == "коридор" && doorlock == true {
		for i := range game.Inventory {
			if game.Inventory[i] == "ключи" {
				doorlock = false
				return "дверь открыта"
			}
		}
	} else if game.State == "коридор" && doorlock == false {
		return "не к чему применить"
	}
	return "не к чему применить"
}

func GameWear(command string) string {
	if command == "рюкзак" && game.Rukzak == false {
	Loop:
		for i := range locations {
			for j := range rooms[locations[i]].RoomLoot {
				if game.State == rooms[locations[i]].RoomName && rooms[locations[i]].RoomLoot[j] == command {
					game.Rukzak = true
					break Loop
				}
			}
		}
		if game.Rukzak == true {
			return "вы надели: рюкзак"
		} else {
			return "нету такого"
		}
	}
	return "нету такого"
}
func GameTake(command string) string {
	var flag bool = false
Loop:
	for i := range locations {
		room := rooms[locations[i]] // берём копию структуры из map
		for j := range room.RoomLoot {
			if game.State == room.RoomName && room.RoomLoot[j] == command {
				// добавляем предмет в инвентарь игрока
				game.Inventory = append(game.Inventory, command)
				flag = true

				// удаляем предмет из комнаты
				room.RoomLoot = append(room.RoomLoot[:j], room.RoomLoot[j+1:]...)

				// сохраняем обратно в map
				rooms[locations[i]] = room

				break Loop
			}
		}
	}
	if flag {
		return "предмет добавлен в инвентарь: " + command
	} else {
		return "нету такого"
	}
}

func handleCommand(command string) string {
	words := strings.Split(command, " ")
	actionshandle := words[0]
	if len(words) == 1 && actionshandle == "осмотреться" {
		return GameLook()
	} else if len(words) == 2 && actionshandle == "идти" {
		return GameGo(words[1])
	} else if len(words) == 2 && actionshandle == "взять" {
		return GameTake(words[1])
	} else if len(words) == 2 && actionshandle == "применить" {
		return GameUse(words[1])
	} else if len(words) == 2 && actionshandle == "надеть" {
		return GameWear(words[1])
	} else {
		return "неизвестная команда"
	}
}

func main() {
	initGame() // инициализация игры

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Добро пожаловать в игру! (напишите 'выход', чтобы завершить)")
	for {
		fmt.Print("> ") // приглашение к вводу
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input) // убираем \n и пробелы

		if input == "выход" {
			fmt.Println("Игра завершена. До встречи!")
			break
		}

		result := handleCommand(input)
		fmt.Println(result)
		//----
	}
}
