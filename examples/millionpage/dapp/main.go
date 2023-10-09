// Copyright (c) Gabriel de Quadros Ligneul
// Copyright (c) Marcel Moura
// SPDX-License-Identifier: MIT (see LICENSE)

// Package that defines the contract logic
package main

import (
	"fmt"
	"image"
	"million"

	"github.com/gligneul/eggroll"
	"github.com/kelindar/bitmap"
)

// DApp contract
type MillionContract struct {
	eggroll.DefaultContract
	Pixels bitmap.Bitmap
}

func (c MillionContract) Clear() {
	c.Pixels.Clear()
}

func (c MillionContract) Paint(p image.Point) error {
	c.Pixels.Set(uint32(p.X + p.Y*1000))
	//fmt.Errorf("invalid coordinates")
	return nil
}

func (c *MillionContract) Decoders() []eggroll.Decoder {
	return []eggroll.Decoder{
		eggroll.NewGenericDecoder[million.Init](),
		eggroll.NewGenericDecoder[million.Paint](),
	}
}

func (c *MillionContract) Advance(env *eggroll.Env, input any) ([]byte, error) {
	switch input := input.(type) {
	case *million.Init:
		c.Clear()
		env.Logf("received input Init. State= %v", c)

	case *million.Paint:
		// TODO Check for valid coordinates. Only free pixels are available
		env.Logf("received input Paint with '%v'\n", input)
		if err := c.Paint(input.Point); err != nil {
			return nil, fmt.Errorf("invalid coordinates")
		}
		env.Logf("DApp state= %v", c)
	default:
		return nil, fmt.Errorf("invalid input")
	}

	env.Logf("c.Pixels.ToBytes()= %v", c.Pixels.ToBytes())
	return []byte("x"), nil // Returning the bitmap because we can
}

func main() {
	eggroll.Roll(&MillionContract{})
}
