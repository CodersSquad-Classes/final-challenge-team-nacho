package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	xSpawnPos, ySpawnPos int
	xPos, yPos           int
	lives                int
	facingDirection      Direction
	sprite               *ebiten.Image
	score                int
}

func (player *Player) initPlayer(xPos int, yPos int, facingDirection Direction) {

	player.xPos = xPos
	player.xSpawnPos = xPos
	player.yPos = yPos
	player.ySpawnPos = yPos
	player.facingDirection = facingDirection

	player.sprite = ebiten.NewImage(CELLSIZE*.75, CELLSIZE*.75)
	player.sprite.Fill(color.RGBA{255, 255, 255, 255})
}

func (player *Player) respawnPlayer() {
	player.xPos = player.xSpawnPos
	player.yPos = player.ySpawnPos
	player.facingDirection = EAST
}

func (player *Player) movePlayer() {
	player.xPos += CardinalDirections[player.facingDirection].xDir
	player.yPos += CardinalDirections[player.facingDirection].yDir
}

func (player *Player) processPlayerInput(input Input, level *World) {

	previousFacingDir := player.facingDirection
	switch input {
	case UP:
		player.facingDirection = NORTH
	case DOWN:
		player.facingDirection = SOUTH
	case LEFT:
		player.facingDirection = WEST
	case RIGHT:
		player.facingDirection = EAST
	}

	newDirection := CardinalDirections[player.facingDirection]

	//Checks if you can move in the new direction, if not, doesn't turn
	if level.levelMatrix[player.yPos+newDirection.yDir][player.xPos+newDirection.xDir].CellType == hasWALL {
		player.facingDirection = previousFacingDir
	}
}

func (player *Player) Update(level *World) {
	for {
		player.processPlayerInput(readInput(), level)
		player.movePlayer()
		level.checkPlayerCollisions()
		time.Sleep(SLEEPTIME)
	}

}

func (player *Player) Draw(screen *ebiten.Image) {

	cellOffset := (CELLSIZE - CELLSIZE*.75) / 2

	opts := &ebiten.DrawImageOptions{}
	xPos := float64(player.xPos*20) + cellOffset + WIDTHOFFSET
	yPos := float64(player.yPos*20) + cellOffset
	opts.GeoM.Translate(float64(xPos), float64(yPos))
	screen.DrawImage(player.sprite, opts)
}
