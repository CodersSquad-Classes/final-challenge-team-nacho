package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Ghost struct {
	xSpawnPos, ySpawnPos int
	xPrev, yPrev         int
	xPos, yPos           int
	facingDirection      Direction
	sprite               *ebiten.Image
	regularSprite        *ebiten.Image
	escapingSprite       *ebiten.Image
	state                GhostState
	isActive             bool
	searchTimer          int
}

//Ghost States
type GhostState int

const (
	WAITING   GhostState = iota
	SEARCHING            = iota
	ESCAPING             = iota
)

func (ghost *Ghost) initGhost(xPos int, yPos int) {
	ghost.isActive = true

	ghost.xPos = xPos
	ghost.xSpawnPos = xPos
	ghost.yPos = yPos
	ghost.ySpawnPos = yPos
	ghost.facingDirection = NORTH

	ghost.regularSprite = ebiten.NewImage(CELLSIZE*.75, CELLSIZE*.75)
	ghost.regularSprite.Fill(color.RGBA{255, 0, 0, 255})
	ghost.escapingSprite = ebiten.NewImage(CELLSIZE*.75, CELLSIZE*.75)
	ghost.escapingSprite.Fill(color.RGBA{255, 0, 255, 255})

	ghost.sprite = ghost.regularSprite
	ghost.state = SEARCHING
}

func (ghost *Ghost) respawnGhost() {
	ghost.xPos = ghost.xSpawnPos
	ghost.yPos = ghost.ySpawnPos
	ghost.facingDirection = NORTH
	ghost.state = SEARCHING
	ghost.sprite = ghost.regularSprite
}

func (ghost *Ghost) moveGhost() {
	ghost.xPrev = ghost.xPos
	ghost.yPrev = ghost.yPos
	ghost.xPos += CardinalDirections[ghost.facingDirection].xDir
	ghost.yPos += CardinalDirections[ghost.facingDirection].yDir
}

func (ghost *Ghost) changeDirection() {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	switch ghost.facingDirection {
	case NORTH:
		fallthrough
	case SOUTH:
		if random.Intn(2) == 0 {
			ghost.facingDirection = EAST
		} else {
			ghost.facingDirection = WEST
		}
	case EAST:
		fallthrough
	case WEST:
		if random.Intn(2) == 0 {
			ghost.facingDirection = NORTH
		} else {
			ghost.facingDirection = SOUTH
		}
	}

}

func (ghost *Ghost) startEscaping() int {
	ghost.state = ESCAPING
	ghost.sprite = ghost.escapingSprite

	switch ghost.facingDirection {
	case NORTH:
		ghost.facingDirection = SOUTH
	case SOUTH:
		ghost.facingDirection = NORTH
	case EAST:
		ghost.facingDirection = WEST
	case WEST:
		ghost.facingDirection = EAST
	}
	return 50
}
func (ghost *Ghost) stopEscaping() {
	ghost.state = SEARCHING
	ghost.sprite = ghost.regularSprite
}

func (ghost *Ghost) runGhost(world *World) {
	for {

		select {
		case <-stopSignal:
			return

		default:
			select {
			case <-escapeSignal:
				ghost.searchTimer = ghost.startEscaping()
			default:
				if ghost.searchTimer == 0 {
					ghost.stopEscaping()
				} else {
					ghost.searchTimer--
				}
			}

			hasCollided := world.checkGhostCollisions(ghost)
			ghost.moveGhost()
			hasCollided = world.checkGhostCollisions(ghost)
			if hasCollided {
				ghost.changeDirection()
			}

			time.Sleep(SLEEPTIME)
		}
	}

}

func (ghost *Ghost) Draw(screen *ebiten.Image) {

	cellOffset := (CELLSIZE - CELLSIZE*.75) / 2

	opts := &ebiten.DrawImageOptions{}
	xPos := float64(ghost.xPos*20) + cellOffset + WIDTHOFFSET
	yPos := float64(ghost.yPos*20) + cellOffset
	opts.GeoM.Translate(float64(xPos), float64(yPos))
	screen.DrawImage(ghost.sprite, opts)
}
