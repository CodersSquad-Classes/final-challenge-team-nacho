package main

type Ghost struct {
	xSpawnPos, ySpawnPos int
	xPos, yPos           int
	lives                int
	facingDirection      Direction
	isAlive              bool
}

//Ghost States
type GhostState int

const (
	WAITING   GhostState = iota
	SEARCHING            = iota
	ESCAPING             = iota
)
