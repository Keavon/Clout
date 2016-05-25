package game

import (
	"fmt"
	"sort"

	"github.com/keavon/clout/pkg/player"
)

type byDamage []player.Player

func (a byDamage) Len() int           { return len(a) }
func (a byDamage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byDamage) Less(i, j int) bool { return a[i].Damage < a[j].Damage }

// Rankings returns the player rankings of the game
// Currently rankings are soley based on least to greatest environmental damage
func (g Game) Rankings() []player.Player {
	fmt.Println(g.Players)
	// Copy game players to an array
	players := []player.Player{}
	players = g.Players
	fmt.Println(players)

	sort.Sort(byDamage(players))

	return players
}
