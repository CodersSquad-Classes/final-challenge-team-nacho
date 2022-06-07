package main

import (
	"flag"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type PacmanGame struct {
	world *World
}

//GAMESTATE consts
type GameState int

const (
	INGAME GameState = iota
	WON              = iota
	LOST             = iota
)

//Global variables
var ghostNum int = 1
var boardX int = 28
var boardY int = 24
var blockSize = 1
var pacmanGame PacmanGame
var stopSignal = make(chan bool)
var escapeSignal = make(chan GhostState)

//Player Functions

// Ghost functions

func setGhostNum(numberOfGhosts int) {
	ghostNum = numberOfGhosts
	if ghostNum > 4 {
		ghostNum = 4
	}
	if ghostNum <= 0 {
		ghostNum = 1
	}
}

// Game implements ebiten.Game interface.
type Game struct{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	pacmanGame.world.Draw(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(WIDTH), HEIGHT
}

func pacmanGameInit(numGhosts int) {
	var world World
	pacmanGame.world = &world
	pacmanGame.world.initWorld(numGhosts)
}

func main() {
	game := &Game{}
	var numGhosts = flag.Int("ghosts", 1, "Number of ghosts in the game, must be in range [1, 4]")
	flag.Parse()

	if *numGhosts < 1 {
		*numGhosts = 1
	} else if *numGhosts > 4 {
		*numGhosts = 4
	}

	pacmanGameInit(*numGhosts)
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(int(WIDTH), HEIGHT)
	ebiten.SetWindowTitle("pacman.go")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
