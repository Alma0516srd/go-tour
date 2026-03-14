package main

import "fmt"

type Player struct {
	name   string
	atBats int
	hits   int
}

func (player *Player) getAverage() float64 {
	return float64(player.hits) / float64(player.atBats)
}

func main() {
	players := []Player{
		{"pl1", 10, 8},
		{"pl2", 44, 5},
		{"pl3", 90, 54},
	}
	for _, player := range players {
		fmt.Println(player.name, player.getAverage())
	}
}
