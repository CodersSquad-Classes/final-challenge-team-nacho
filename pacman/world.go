package main

import (
	"bufio"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type CellType int

const (
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
var ghost0, ghost1, ghost2, ghost3, ghost4 Ghost

type World struct {
	player       *Player
	levelMatrix  [][]Cell
	ghosts       [4]*Ghost
	levelPellets int
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

	var ghostCounter = 0

	//Thios is not ideal, but due to time constraints was the only way I could think of doing it
	for i := range world.ghosts {
		switch i {
		case 0:
			world.ghosts[i] = &ghost0
		case 1:
			world.ghosts[i] = &ghost1
		case 2:
			world.ghosts[i] = &ghost2
		case 3:
			world.ghosts[i] = &ghost3
		case 4:
			world.ghosts[i] = &ghost4
		}
		world.ghosts[i].isActive = false
	}

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
				world.ghosts[ghostCounter].initGhost(xPos, yPos)
				ghostCounter++
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

func (world *World) initWorld() {
	world.loadMaze("level.txt")
	go world.player.runPlayer(world)
	for _, ghost := range world.ghosts {
		if ghost.isActive {
			go ghost.runGhost(world)
		}
	}

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

func (world *World) checkGhostCollisions(ghost *Ghost) bool {
	switch world.levelMatrix[ghost.yPos][ghost.xPos].CellType {
	case hasWALL:
		ghost.xPos -= CardinalDirections[ghost.facingDirection].xDir
		ghost.yPos -= CardinalDirections[ghost.facingDirection].yDir
		return true
	}
	if world.player.xPos == ghost.xPos && world.player.yPos == ghost.yPos {
		if ghost.state == SEARCHING {
			world.playerHit()
			return false
		} else {
			ghost.respawnGhost()
		}
	}

	//Prevents the player and a ghost "jumping" over each other if they switch places
	if world.player.xPos == ghost.xPrev && world.player.yPos == ghost.yPrev && world.player.xPrev == ghost.xPos && world.player.yPrev == ghost.yPos {
		if ghost.state == SEARCHING {
			world.playerHit()
			return false
		} else {
			ghost.respawnGhost()
		}

	}
	return false
}

func (world *World) playerHit() {
	if player.lives > 0 {
		player.respawnPlayer()
		for _, ghost := range world.ghosts {
			ghost.respawnGhost()
		}
	} else {
		os.Exit(0)
	}

}

func (world *World) Draw(screen *ebiten.Image) {

	for y := range world.levelMatrix {
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
	for _, ghost := range world.ghosts {
		if ghost.isActive {
			ghost.Draw(screen)
		}
	}
}
