package pacman

import (
	"bufio"
	"log"
	"os"
	"os/exec"
)

//Input Consts
type Input byte

const (
	UP     Input = 1 << iota // 1
	DOWN         = 1 << iota // 2
	LEFT         = 1 << iota // 4
	RIGHT        = 1 << iota // 8
	PAUSE        = 1 << iota // 16
	ESCAPE       = 1 << iota // 32
)

//Input Consts
type Direction byte

const (
	UpDir    Direction = iota // 1
	DownDir            = iota // 2
	LeftDir            = iota // 4
	RightDir           = iota // 8
)

//GAMESTATE consts
type GameState int64

const (
	MENU     GameState = iota
	INGAME             = iota
	PAUSED             = iota
	GAMEOVER           = iota
)

//Ghost States

//Game Objects
type Player struct {
	xPos, yPos      int64
	lives           int64
	facingDirection Direction
	isAlive         bool
}

type Ghost struct {
	xPos, yPos      int64
	lives           int64
	facingDirection Direction
	isAlive         bool
}

//Global variables
var input Input
var gameState GameState = MENU
var level []string

//Player Functions

// Ghost functions

// Game Functions
func initTerminal() {

	//Puts terminal in cbreak mode to register key presses as raw Input
	cbTerm := exec.Command("stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin

	err := cbTerm.Run()
	if err != nil {
		log.Fatalln("unable to activate cbreak mode:", err)
	}
}

func cleanupTerminal() {

	//Restores terminal to cooked mode as it is the default
	cookedTerm := exec.Command("stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin

	err := cookedTerm.Run()
	if err != nil {
		log.Fatalln("unable to restore cooked mode:", err)
	}
}

func loadLevel(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		level = append(level, line)
	}
	return nil
}

func readInput() {

}

func main() {

	//Initialise game
	initTerminal()
	defer cleanupTerminal()

	//Full Loop
	for {
		switch gameState {
		case MENU:
			break
		case INGAME:
			break
		case PAUSE:
			break
		case GAMEOVER:
			os.Exit(0)
		}

	}

}
