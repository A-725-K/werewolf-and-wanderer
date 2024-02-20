package main

import (
	"fmt"
	"werewolves-and-wanderer/m/v2/internal/game"
)

func main() {
	fmt.Print(game.BANNER)
	game.InitGame().Start()
}
