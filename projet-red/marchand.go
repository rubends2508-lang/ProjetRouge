package main

import "fmt"

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
		fmt.Println("7) Forgeron (fabrique avec 2 fer)")
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

func openForgeron(c *Character) {
	if c.Iron < 2 {
		fmt.Println("Pas assez de fer (il faut 2).")
		return
	}
	fmt.Println("\n--- Forgeron ---")
	fmt.Println("1) Forger Épée (+100 dégâts)")
	fmt.Println("2) Forger Plastron (+50 PV max)")
	fmt.Println("3) Forger Jambières (+30 PV max)")
	fmt.Println("4) Forger Casque (+20 PV max)")
	fmt.Println("5) Retour")
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
	case "5":
		return
	}
}
