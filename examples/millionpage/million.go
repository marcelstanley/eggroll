// Copyright (c) Gabriel de Quadros Ligneul
// Copyright (c) Marcel Moura
// SPDX-License-Identifier: MIT (see LICENSE)

// Package million exports shared types for the DApp
package million

import (
	"fmt"
	"image"
	"image/color"

	"github.com/gligneul/eggroll"
	"github.com/kelindar/bitmap"
)

type (
	// Input type Init forces the initialization of the DApp state
	Init struct {
	}

	// Input type Paint specifies a pixel to be painted
	Paint struct {
		Point image.Point
		Color color.RGBA
	}
)

// DApp contract
type Contract struct {
	Pixels bitmap.Bitmap
}

func (c Contract) Clear() {
	c.Pixels.Clear()
}

func (c Contract) Paint(p image.Point) error {
	c.Pixels.Set(uint32(p.X + p.Y*1000))
	return nil
}

func (c *Contract) Decoders() []eggroll.Decoder {
	return []eggroll.Decoder{
		eggroll.NewGenericDecoder[Init](),
		eggroll.NewGenericDecoder[Paint](),
	}
}

func (c *Contract) Advance(env *eggroll.Env, input any) error {
	switch input := input.(type) {
	case *Init:
		c.Clear()
		env.Logf("received input Init. State= %v", c)
	case *Paint:
		// TODO Check for valid coordinates. Only free pixels are available
		env.Logf("received input Paint with '%v'\n", input)
		if c.Paint(input.Point) != nil {
			return fmt.Errorf("invalid coordinates")
		}
	default:
		return fmt.Errorf("invalid input")
	}
	return nil
}
