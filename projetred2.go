package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const MAP_SIZE = 7          // Taille de la carte en (7x7)
const INVENTORY_LIMIT = 10  // la fameuse limite d'objets de l'inventaire a 10

//  la stucture du jeu 

// Objet (Potion, Pancake)
type Item struct {
	Name string
}

// Équipements que le joueur posséde 
type Equipment struct {
	Weapon   string
	Plastron string
	Jambiere string
	Casque   string
}

// Personnage joueur
type Character struct {
	Name      string
	Race      string
	HP        int
	HPMax     int
	Damage    int
	Gems      int
	Keys      int
	Inventory []Item
	Equip     Equipment
	X, Y      int
	Iron      int
}

// Ennemi gobelin ou Boss de fin 
type Monster struct {
	Name   string
	HP     int
	HPMax  int
	Attack int
}


var world [MAP_SIZE][MAP_SIZE]rune
var reader = bufio.NewReader(os.Stdin)

//	Utilitaires 

// Lecture du clavier
func input(prompt string) string {
	fmt.Print(prompt)
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

// Si vide -> affiche (aucun)
func emptyIfNone(s string) string {
	if s == "" {
		return "(aucun)"
	}
	return s
}

// Création du personnage 
// Choix de la race des persos
func chooseRace() (string, int, int) {
	for {
		fmt.Println("Choisis ta race :")
		fmt.Println("1) Mini P.E.K.K.A (100 PV / 75 dégâts)")
		fmt.Println("2) Chevalier (200 PV / 50 dégâts)")
		fmt.Println("3) Barbares (150 PV / 60 dégâts)")
		choice := input("> ")
		switch choice {
		case "1":
			return "Mini P.E.K.K.A", 100, 75
		case "2":
			return "Chevalier", 200, 50
		case "3":
			return "Barbares", 150, 60
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

// Création du joueur complet
func createCharacter() *Character {
	name := input("Entre ton nom : ")
	race, hpmax, dmg := chooseRace()
	return &Character{
		Name:      name,
		Race:      race,
		HPMax:     hpmax,
		HP:        hpmax,
		Damage:    dmg,
		Gems:      100,
		Keys:      0,
		Inventory: []Item{},
		Equip:     Equipment{},
		X:         0,
		Y:         0,
		Iron:      0,
	}
}

//  inventaire de la personne

// Ajouter un objet
func addToInventory(c *Character, it Item) bool {
	if len(c.Inventory) >= INVENTORY_LIMIT {
		fmt.Println("❌ Inventaire plein !")
		return false
	}
	c.Inventory = append(c.Inventory, it)
	return true
}

// Retirer un objet
func removeFromInventory(c *Character, idx int) {
	if idx >= 0 && idx < len(c.Inventory) {
		c.Inventory = append(c.Inventory[:idx], c.Inventory[idx+1:]...)
	}
}

// Afficher l inventaire et utiliser des objets
func showInventory(c *Character, inCombat bool) string {
	fmt.Println("\n--- INVENTAIRE ---")
	fmt.Printf("Nom: %s | Race: %s\n", c.Name, c.Race)
	fmt.Printf("PV: %d/%d | Gemmes: %d | Clés: %d | Fer: %d\n",
		c.HP, c.HPMax, c.Gems, c.Keys, c.Iron)
	fmt.Printf("Équipement: Arme=%s Plastron=%s Jambières=%s Casque=%s\n",
		emptyIfNone(c.Equip.Weapon), emptyIfNone(c.Equip.Plastron),
		emptyIfNone(c.Equip.Jambiere), emptyIfNone(c.Equip.Casque))

	if len(c.Inventory) == 0 {
		fmt.Println("Inventaire vide.")
		return ""
	}
	for i, it := range c.Inventory {
		fmt.Printf("%d) %s\n", i+1, it.Name)
	}

	choice := input("Choisis un objet à utiliser (numéro) ou Enter pour quitter > ")
	if choice == "" {
		return ""
	}
	var idx int
	_, err := fmt.Sscanf(choice, "%d", &idx)
	if err != nil || idx <= 0 || idx > len(c.Inventory) {
		fmt.Println("Choix invalide.")
		return ""
	}
	item := c.Inventory[idx-1]
	switch item.Name {
	case "Pancake":
		c.HP += 50
		if c.HP > c.HPMax {
			c.HP = c.HPMax
		}
		fmt.Printf("🥞 Tu manges un Pancake. PV:%d/%d\n", c.HP, c.HPMax)
		removeFromInventory(c, idx-1)
		return "heal"
	case "Potion de poison":
		if inCombat {
			fmt.Println("☠️ Tu prépares une Potion de poison.")
			removeFromInventory(c, idx-1)
			return "poison"
		} else {
			fmt.Println("La Potion de poison ne sert qu'en combat.")
		}
	}
	return ""
}

//  Carte 

// Génération de la carte avec marchands coffres combats et boss
func generateMap() {
	for y := 0; y < MAP_SIZE; y++ {
		for x := 0; x < MAP_SIZE; x++ {
			world[y][x] = '.'
		}
	}
	placeRandom('M', 3)  // marchands
	placeRandom('C', 10) // combats
	placeRandom('T', 5)  // coffres
	world[MAP_SIZE-1][MAP_SIZE-1] = 'B' // boss final
}

// Place des éléments aléatoirement
func placeRandom(symbol rune, count int) {
	for placed := 0; placed < count; {
		x := rand.Intn(MAP_SIZE)
		y := rand.Intn(MAP_SIZE)
		if world[y][x] == '.' && !(x == 0 && y == 0) {
			world[y][x] = symbol
			placed++
		}
	}
}

// Affiche la carte et la position du joueur
func printMap(c *Character) {
	fmt.Println("\nCarte (P=toi, M=marchand, C=combats, T=coffre, B=boss)")
	for y := 0; y < MAP_SIZE; y++ {
		for x := 0; x < MAP_SIZE; x++ {
			if c.X == x && c.Y == y {
				fmt.Print("P ")
			} else {
				fmt.Printf("%c ", world[y][x])
			}
		}
		fmt.Println()
	}
}

//  Marchand et Forgeron 

// Menu du marchand
func openMerchant(c *Character) {
	for {
		fmt.Println("\n--- Marchand ---")
		fmt.Printf("Gemmes: %d | Fer: %d\n", c.Gems, c.Iron)
		fmt.Println("1) Pancake (10 gemmes)")
		fmt.Println("2) Potion de poison (15 gemmes)")
		fmt.Println("3) Épée (55 gemmes)")
		fmt.Println("4) Plastron (30 gemmes)")
		fmt.Println("5) Jambières (20 gemmes)")
		fmt.Println("6) Casque (25 gemmes)")
		fmt.Println("7) Forgeron (utilise 2 fer)")
		fmt.Println("8) Quitter")
		choice := input("> ")
		switch choice {
		case "1":
			if c.Gems >= 10 && addToInventory(c, Item{"Pancake"}) {
				c.Gems -= 10
				fmt.Println("🥞 Pancake acheté.")
			}
		case "2":
			if c.Gems >= 15 && addToInventory(c, Item{"Potion de poison"}) {
				c.Gems -= 15
				fmt.Println("☠️ Potion de poison achetée.")
			}
		case "3":
			if c.Gems >= 55 {
				c.Gems -= 55
				c.Equip.Weapon = "Épée"
				c.Damage += 100
				fmt.Println("⚔️ Tu as équipé une Épée (+100 dégâts).")
			}
		case "4":
			if c.Gems >= 30 {
				c.Gems -= 30
				c.Equip.Plastron = "Plastron"
				c.HPMax += 50
				fmt.Println("🛡️ Plastron équipé (+50 PV max).")
			}
		case "5":
			if c.Gems >= 20 {
				c.Gems -= 20
				c.Equip.Jambiere = "Jambières"
				c.HPMax += 30
				fmt.Println("👖 Jambières équipées (+30 PV max).")
			}
		case "6":
			if c.Gems >= 25 {
				c.Gems -= 25
				c.Equip.Casque = "Casque"
				c.HPMax += 20
				fmt.Println("🪖 Casque équipé (+20 PV max).")
			}
		case "7":
			openForgeron(c)
		case "8":
			return
		}
	}
}

// Menu du forgeron
func openForgeron(c *Character) {
	if c.Iron < 2 {
		fmt.Println("Pas assez de fer (2 nécessaires).")
		return
	}
	fmt.Println("1) Forger Épée (+100 dégâts)")
	fmt.Println("2) Forger Plastron (+50 PV max)")
	fmt.Println("3) Forger Jambières (+30 PV max)")
	fmt.Println("4) Forger Casque (+20 PV max)")
	choice := input("> ")
	switch choice {
	case "1":
		c.Iron -= 2
		c.Equip.Weapon = "Épée forgée"
		c.Damage += 100
		fmt.Println("⚒️ Épée forgée (+100 dégâts).")
	case "2":
		c.Iron -= 2
		c.Equip.Plastron = "Plastron forgé"
		c.HPMax += 50
		fmt.Println("⚒️ Plastron forgé (+50 PV max).")
	case "3":
		c.Iron -= 2
		c.Equip.Jambiere = "Jambières forgées"
		c.HPMax += 30
		fmt.Println("⚒️ Jambières forgées (+30 PV max).")
	case "4":
		c.Iron -= 2
		c.Equip.Casque = "Casque forgé"
		c.HPMax += 20
		fmt.Println("⚒️ Casque forgé (+20 PV max).")
	}
}

//  Coffres 
func openTreasure(c *Character, x, y int) {
	if c.Keys <= 0 {
		fmt.Println("🔑 Il faut une clé pour ouvrir ce coffre !")
		return
	}
	c.Keys--
	if rand.Intn(2) == 0 {
		c.Iron++
		fmt.Println("🎁 Tu trouves 1 fer.")
	} else {
		c.Gems += 25
		fmt.Println("🎁 Tu trouves 25 gemmes.")
	}
	world[y][x] = '.'
}

//  Combat 
func combat(c *Character, boss bool) {
	var m Monster
	if boss {
		m = Monster{"Gobelinstein", 1000, 1000, 80}
	} else {
		m = Monster{"Gobelin", 200, 200, 40}
	}
	for c.HP > 0 && m.HP > 0 {
		fmt.Printf("\n%s PV:%d/%d | %s PV:%d/%d\n",
			c.Name, c.HP, c.HPMax, m.Name, m.HP, m.HPMax)
		fmt.Println("1) Coup de poing")
		fmt.Println("2) Attaquer avec arme")
		fmt.Println("3) Inventaire")
		fmt.Println("4) Fuir (sauf Boss)")
		choice := input("> ")

		if choice == "1" {
			m.HP -= c.Damage
		} else if choice == "2" {
			if c.Equip.Weapon != "" {
				m.HP -= c.Damage
			} else {
				fmt.Println("Pas d'arme équipée.")
			}
		} else if choice == "3" {
			effect := showInventory(c, true)
			if effect == "poison" {
				for i := 0; i < 3; i++ {
					time.Sleep(1 * time.Second)
					m.HP -= 10
					fmt.Printf("☠️ Poison... %s PV:%d/%d\n", m.Name, m.HP, m.HPMax)
				}
			}
		} else if choice == "4" && !boss {
			fmt.Println("Tu prends la fuite...")
			return
		}

		if m.HP <= 0 {
			fmt.Println("✅ Tu as vaincu ", m.Name)
			if boss {
				fmt.Println("🎉 Félicitations, tu as terminé PROJET RED ! 🎉")
				os.Exit(0)
			} else {
				c.Gems += 25
				c.Keys++
				fmt.Println("Butin : 25 gemmes + 1 clé.")
			}
			return
		}

		// Tour du monstre
		c.HP -= m.Attack
		fmt.Printf("%s t'attaque (-%d PV)\n", m.Name, m.Attack)
		if c.HP <= 0 {
			fmt.Println("💀 Tu es mort... Game Over.")
			os.Exit(0)
		}
	}
}

// Menu principal
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

// Point d’entrée 
func main() {
	mainMenu()
}
