package main

//Input Consts
type Direction int

const (
	NORTH Direction = iota // 0
	SOUTH           = iota // 1
	WEST            = iota // 2
	EAST            = iota // 3
)

type DirPair struct {
	xDir, yDir int
}

var CardinalDirections [4]DirPair = [4]DirPair{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
