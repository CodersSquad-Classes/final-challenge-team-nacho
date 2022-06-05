package main

import "github.com/hajimehoshi/ebiten/v2"

//Input Consts
type Input byte

const (
	UP     Input = 1 << iota // 1
	DOWN         = 1 << iota // 2
	LEFT         = 1 << iota // 4
	RIGHT        = 1 << iota // 8
	ESCAPE       = 1 << iota // 16
	START        = 1 << iota // 32
	NONE         = 1 << iota // 64
)

func readInput() Input {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		return UP
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		return DOWN
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		return LEFT
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		return RIGHT
	}
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ESCAPE
	}
	return NONE
}
