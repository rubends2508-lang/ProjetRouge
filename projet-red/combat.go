package main

import (
	"fmt"
	"os"
	"time"
)

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
			fmt.Println("Tu frappes (", c.Damage, " dégâts)")
		} else if choice == "2" {
			if c.Equip.Weapon != "" {
				m.HP -= c.Damage
				fmt.Println("Tu attaques avec ton arme (", c.Damage, " dégâts)")
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
					if m.HP <= 0 {
						break
					}
				}
			}
		} else if choice == "4" {
			if boss {
				fmt.Println("Impossible de fuir le Boss !")
			} else {
				fmt.Println("Tu prends la fuite...")
				return
			}
		}

		if m.HP <= 0 {
			fmt.Println("✅ Tu as vaincu ", m.Name)
			if boss {
				fmt.Println("🎉 Félicitations, tu as vaincu le Gobelinstein et terminé PROJET RED ! 🎉")
				os.Exit(0)
			} else {
				c.Gems += 25
				c.Keys++
				fmt.Println("Tu gagnes 25 gemmes et une clé !")
			}
			return
		}

		fmt.Printf("%s t'attaque (-%d PV)\n", m.Name, m.Attack)
		c.HP -= m.Attack
		if c.HP <= 0 {
			fmt.Println("💀 Tu es mort... Game Over.")
			os.Exit(0)
		}
	}
}
