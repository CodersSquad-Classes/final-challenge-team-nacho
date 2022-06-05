package main

import "time"

const (
	WIDTH       float64       = 640
	HEIGHT                    = 560
	CELLSIZE                  = 20
	WIDTHOFFSET               = (WIDTH - (28 * CELLSIZE)) / 2
	SLEEPTIME   time.Duration = time.Second / 5
)
