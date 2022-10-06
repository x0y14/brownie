package brownie

import (
	"golang.org/x/image/font/sfnt"
)

type Style struct {
	fontSize int
	font     *sfnt.Font
}
