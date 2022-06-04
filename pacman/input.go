package pacman

//Input Consts
type Input byte

const (
	UP     Input = 1 << iota // 1
	DOWN         = 1 << iota // 2
	LEFT         = 1 << iota // 4
	RIGHT        = 1 << iota // 8
	ESCAPE       = 1 << iota // 16
	START        = 1 << iota // 32
	NONE         = 1 << iota // 64
)
