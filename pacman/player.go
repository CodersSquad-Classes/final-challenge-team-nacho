package pacman

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	xPos, yPos      int
	lives           int
	facingDirection Direction
}

func (player *Player) Draw(screen *ebiten.Image) {
	var image *ebiten.Image
	image = ebiten.NewImage(CELLSIZE*.75, CELLSIZE*.75)
	image.Fill(color.RGBA{255, 255, 255, 255})
	cellOffset := (CELLSIZE - CELLSIZE*.75) / 2

	opts := &ebiten.DrawImageOptions{}
	xPos := float64(player.xPos*20) + cellOffset + WIDTHOFFSET
	yPos := float64(player.yPos*20) + cellOffset
	opts.GeoM.Translate(float64(xPos), float64(yPos))
	screen.DrawImage(image, opts)
}
