// Copyright (c) Marcel Moura
// SPDX-License-Identifier: MIT (see LICENSE)

// Package million exports shared types for the DApp
package million

import (
	"image"
	"image/color"
)

// XXX Remove
// Input type Init forces the initialization of the DApp state
type Init struct {
}

// Input type Paint specifies a pixel to be painted
type Paint struct {
	Point image.Point
	Color color.RGBA
}
