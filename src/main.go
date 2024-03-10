package main

import (
	"project/src/game"
)

func main() {
	game := game.NewGame()
	game.Init()
	game.Run()
	game.Close()
}
