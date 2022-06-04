package pacman

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

var WALLCELL, EMPTYCELL, PELLETCELL, POWERCELL, PLAYERCELL *ebiten.Image
var player Player

type Level struct {
	player       *Player
	levelMatrix  [][]Cell
	levelPellets int
	//activeGhosts []Ghost
}

func (level *Level) loadMaze(file string) error {
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

	level.levelMatrix = make([][]Cell, len(maze))
	for i := range level.levelMatrix {
		level.levelMatrix[i] = make([]Cell, len(maze[i]))
	}

	level.levelMatrix = make([][]Cell, len(maze))
	for i := range level.levelMatrix {
		level.levelMatrix[i] = make([]Cell, len(maze[i]))
		for j := range level.levelMatrix[i] {
			level.levelMatrix[i][j].x = j
			level.levelMatrix[i][j].y = i
			level.levelMatrix[i][j].CellType = EMPTY
		}
	}

	level.levelPellets = 0
	level.player = &player

	for yPos := range maze {
		for xPos, cell := range maze[yPos] {
			switch cell {
			case '#':
				level.levelMatrix[yPos][xPos].CellType = hasWALL
			case '.':
				level.levelMatrix[yPos][xPos].CellType = hasPELLET
				level.levelPellets++
			case 'X':
				level.levelMatrix[yPos][xPos].CellType = hasPOWER
			case 'P':
				level.player.xPos = xPos
				level.player.yPos = yPos
				level.player.lives = 3
				level.player.facingDirection = EAST
			case 'G':
				fallthrough
			case ' ':
				level.levelMatrix[yPos][xPos].CellType = EMPTY
			}

		}
	}
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
func (level *Level) Draw(screen *ebiten.Image) {

	for y, _ := range level.levelMatrix {
		for x, cell := range level.levelMatrix[y] {
			var image *ebiten.Image
			var cellOffset float64 = 0

			switch cell.CellType {
			case EMPTY:
				image = ebiten.NewImage(CELLSIZE, CELLSIZE)
				image.Fill(color.Black)
			case hasWALL:
				image = ebiten.NewImage(CELLSIZE, CELLSIZE)
				image.Fill(color.RGBA{0, 0, 255, 255})
			case hasPELLET:
				image = ebiten.NewImage(CELLSIZE*.25, CELLSIZE*.25)
				image.Fill(color.RGBA{255, 255, 0, 255})
				cellOffset = (CELLSIZE - CELLSIZE*.25) / 2
			case hasPOWER:
				image = ebiten.NewImage(CELLSIZE*.5, CELLSIZE*.5)
				image.Fill(color.RGBA{0, 255, 255, 255})
				cellOffset = (CELLSIZE - CELLSIZE*.5) / 2

			}

			opts := &ebiten.DrawImageOptions{}
			xPos := float64(x*20) + cellOffset + WIDTHOFFSET
			yPos := float64(y*20) + cellOffset
			opts.GeoM.Translate(float64(xPos), float64(yPos))
			screen.DrawImage(image, opts)
		}
	}
	level.player.Draw(screen)
}
