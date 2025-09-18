package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	mainMenu()
}

//  Menu Principal 
func mainMenu() {
	for {
		fmt.Println("\n=== MENU PRINCIPAL ===")
		fmt.Println("1) Nouvelle partie")
		fmt.Println("2) Qui sont-ils ?")
		fmt.Println("3) Quitter")
		choice := input("> ")
		if choice == "1" {
			runGame()
		} else if choice == "2" {
			fmt.Println("🎭 Les artistes cachés sont : Ruben et Kenley Lebogos")
		} else if choice == "3" {
			fmt.Println("Au revoir.")
			os.Exit(0)
		}
	}
}

//  Jeu principal 
func runGame() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(`
██████╗ ██████╗  ██████╗      ██╗     ███████╗██████╗ 
██╔══██╗██╔══██╗██╔═══██╗     ██║     ██╔════╝██╔══██╗
██████╔╝██████╔╝██║   ██║     ██║     █████╗  ██████╔╝
██╔═══╝ ██╔═══╝ ██║   ██║     ██║     ██╔══╝  ██╔═══╝ 
██║     ██║     ╚██████╔╝     ███████╗███████╗██║     
╚═╝     ╚═╝      ╚═════╝      ╚══════╝╚══════╝╚═╝     
                >>> PROJET RED <<<
	`)

	player := createCharacter()
	generateMap()

	for {
		printMap(player)
		fmt.Printf("Position (%d,%d) | PV:%d/%d | Gemmes:%d | Clés:%d | Fer:%d | Inventaire:%d/%d\n",
			player.X, player.Y, player.HP, player.HPMax,
			player.Gems, player.Keys, player.Iron,
			len(player.Inventory), INVENTORY_LIMIT)

		fmt.Println("Déplacements : Z/S/Q/D | I=Inventaire | Quitter")
		cmd := strings.ToUpper(input("> "))
		switch cmd {
		case "Z":
			if player.Y > 0 {
				player.Y--
			}
		case "S":
			if player.Y < MAP_SIZE-1 {
				player.Y++
			}
		case "Q":
			if player.X > 0 {
				player.X--
			}
		case "D":
			if player.X < MAP_SIZE-1 {
				player.X++
			}
		case "I":
			showInventory(player, false)
			continue
		case "QUITTER":
			fmt.Println("Fin du jeu.")
			return
		}
		cell := world[player.Y][player.X]
		if cell == 'M' {
			openMerchant(player)
		}
		if cell == 'C' {
			combat(player, false)
		}
		if cell == 'T' {
			openTreasure(player, player.X, player.Y)
		}
		if cell == 'B' {
			combat(player, true)
		}
	}
}
