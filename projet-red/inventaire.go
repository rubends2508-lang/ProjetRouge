package main

import (
	"fmt"
)

const INVENTORY_LIMIT = 10

func addToInventory(c *Character, it Item) bool {
	if len(c.Inventory) >= INVENTORY_LIMIT {
		fmt.Println("❌ Inventaire plein !")
		return false
	}
	c.Inventory = append(c.Inventory, it)
	return true
}

func removeFromInventory(c *Character, idx int) {
	if idx >= 0 && idx < len(c.Inventory) {
		c.Inventory = append(c.Inventory[:idx], c.Inventory[idx+1:]...)
	}
}

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

func emptyIfNone(s string) string {
	if s == "" {
		return "(aucun)"
	}
	return s
}
