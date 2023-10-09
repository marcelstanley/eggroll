// Copyright (c) Marcel Moura
// SPDX-License-Identifier: MIT (see LICENSE)

// Package million exports shared types for the DApp
package million

import (
	"image"
	"image/color"
)

const (
	MAX_VALUE = 999
	SUCCESS   = "S"
	FAILURE   = "F"
)

// Input type Paint specifies a pixel to be painted
type Paint struct {
	Point image.Point
	Color color.RGBA
}
