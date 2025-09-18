package main

type Item struct {
	Name string
}

type Equipment struct {
	Weapon   string
	Plastron string
	Jambiere string
	Casque   string
}

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

type Monster struct {
	Name   string
	HP     int
	HPMax  int
	Attack int
}
