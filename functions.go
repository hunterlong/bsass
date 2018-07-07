package bsass

import (
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
)

func darken(hex string, percent float64) string {
	c, _ := colorful.Hex(hex)
	r, g, b := c.RGB255()
	lessR := float64(r) * (percent * 0.01) / 255
	lessG := float64(g) * (percent * 0.01) / 255
	lessB := float64(b) * (percent * 0.01) / 255
	newColor := colorful.Color{float64(lessR), float64(lessG), float64(lessB)}
	return newColor.Hex()
}

func lighten(hex string, percent float64) string {
	c, _ := colorful.Hex(hex)
	r, g, b := c.RGB255()
	lessR := float64(r) * (percent * 0.01) / 255
	lessG := float64(g) * (percent * 0.01) / 255
	lessB := float64(b) * (percent * 0.01) / 255

	fmt.Println(float64(r) * (percent * 0.01))
	fmt.Println(r, lessR)

	newColor := colorful.Color{float64(lessR), float64(lessG), float64(lessB)}
	return newColor.Hex()
}
