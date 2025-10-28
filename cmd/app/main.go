package main

import "module1/internal/domain"

func main() {
	game := domain.NewGame()
	game.GameCycle()
}
