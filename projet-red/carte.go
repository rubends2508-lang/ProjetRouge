package main

import (
	"fmt"
	"math/rand"
)

const MAP_SIZE = 7

var world [MAP_SIZE][MAP_SIZE]rune

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
