package main

import (
	"bufio"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type CellType int

const (
	hasPLAYER = iota
	hasGHOST  = iota
	hasWALL   = iota
	hasPELLET = iota
	hasPOWER  = iota
	EMPTY     = iota
)

type Cell struct {
	x, y     int
	CellType CellType
}

var WALLCELL, EMPTYCELL, PELLETCELL, POWERCELL *ebiten.Image
var player Player

type World struct {
	player       *Player
	levelMatrix  [][]Cell
	levelPellets int
	//activeGhosts []Ghost
}

func initCells() {
	WALLCELL = ebiten.NewImage(CELLSIZE, CELLSIZE)
	WALLCELL.Fill(color.RGBA{0, 0, 255, 255})
	EMPTYCELL = ebiten.NewImage(CELLSIZE, CELLSIZE)
	EMPTYCELL.Fill(color.Black)
	PELLETCELL = ebiten.NewImage(CELLSIZE*.25, CELLSIZE*.25)
	PELLETCELL.Fill(color.RGBA{255, 255, 0, 255})
	POWERCELL = ebiten.NewImage(CELLSIZE*.5, CELLSIZE*.5)
	POWERCELL.Fill(color.RGBA{0, 255, 255, 255})
}

func (world *World) loadMaze(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineIndx := 0
	var maze [][]byte
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, []byte(line))
		lineIndx++
	}

	world.levelMatrix = make([][]Cell, len(maze))
	for i := range world.levelMatrix {
		world.levelMatrix[i] = make([]Cell, len(maze[i]))
	}

	world.levelMatrix = make([][]Cell, len(maze))
	for i := range world.levelMatrix {
		world.levelMatrix[i] = make([]Cell, len(maze[i]))
		for j := range world.levelMatrix[i] {
			world.levelMatrix[i][j].x = j
			world.levelMatrix[i][j].y = i
			world.levelMatrix[i][j].CellType = EMPTY
		}
	}

	world.levelPellets = 0
	world.player = &player

	for yPos := range maze {
		for xPos, cell := range maze[yPos] {
			switch cell {
			case '#':
				world.levelMatrix[yPos][xPos].CellType = hasWALL
			case '.':
				world.levelMatrix[yPos][xPos].CellType = hasPELLET
				world.levelPellets++
			case 'X':
				world.levelMatrix[yPos][xPos].CellType = hasPOWER
			case 'P':
				world.player.initPlayer(xPos, yPos, EAST)
			case 'G':
				fallthrough
			case ' ':
				world.levelMatrix[yPos][xPos].CellType = EMPTY
			}

		}
	}
	initCells()
	return nil
}

/*func (level Level) updateLevelDsiplay() {
	for x, _ := range level.levelMatrix {
		for y, _ := range level.levelMatrix[x] {
			switch level.levelMatrix[x][y] {
			case hasPLAYER:
				level.levelDisplay[x][y] = PLAYERCELL
			case hasWALL:
				level.levelDisplay[x][y] = WALLCELL
			case hasPOWER:
				level.levelDisplay[x][y] = POWERCELL
			case hasPELLET:
				level.levelDisplay[x][y] = PELLETCELL
			case EMPTY:
				level.levelDisplay[x][y] = EMPTYCELL
			}

		}
	}
}
*/

func (world *World) initLevel() {
	world.loadMaze("level.txt")
	go world.player.Update(world)

}

func (world *World) checkPlayerCollisions() {
	switch world.levelMatrix[world.player.yPos][world.player.xPos].CellType {
	case hasWALL:
		world.player.xPos -= CardinalDirections[world.player.facingDirection].xDir
		world.player.yPos -= CardinalDirections[world.player.facingDirection].yDir
	case hasPOWER:
		fallthrough
	case hasPELLET:
		world.levelMatrix[world.player.yPos][world.player.xPos].CellType = EMPTY
		world.player.score++
	}
}

func (world *World) Draw(screen *ebiten.Image) {

	for y, _ := range world.levelMatrix {
		for x, cell := range world.levelMatrix[y] {
			var image *ebiten.Image
			var cellOffset float64 = 0

			switch cell.CellType {
			case EMPTY:
				image = EMPTYCELL
			case hasWALL:
				image = WALLCELL
			case hasPELLET:
				image = PELLETCELL
				cellOffset = (CELLSIZE - CELLSIZE*.25) / 2
			case hasPOWER:
				image = POWERCELL
				cellOffset = (CELLSIZE - CELLSIZE*.5) / 2

			}

			opts := &ebiten.DrawImageOptions{}
			xPos := float64(x*20) + cellOffset + WIDTHOFFSET
			yPos := float64(y*20) + cellOffset
			opts.GeoM.Translate(float64(xPos), float64(yPos))
			screen.DrawImage(image, opts)
		}
	}
	world.player.Draw(screen)
}
