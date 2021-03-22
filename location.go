package battleshipgameng

import (
	"fmt"
	"strings"
)

const ascii string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type location struct {
	y   int
	x   int
	hit bool
}

func locToString(y, x int) string {
	return fmt.Sprintf("%v%v", string(ascii[y]), x)
}

func locToInt(loc string) (y, x int) {
	y = strings.IndexByte(ascii, loc[y])
	x = int(loc[1])

	return y, x
}

func (loc location) getStringLoc() string {
	return fmt.Sprintf("%v%v", string(ascii[loc.y]), loc.x)
}
