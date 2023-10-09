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

const (
	MAX     = 999
	SUCCESS = "S"
	FAILURE = "F"
)

// DApp contract
type MillionContract struct {
	eggroll.DefaultContract
	Pixels bitmap.Bitmap
}

func (c *MillionContract) Clear() {
	c.Pixels.Clear()
}

func (c *MillionContract) Paint(p image.Point) error {
	if p.X < 0 || p.Y < 0 || p.X > MAX || p.Y > MAX {
		return fmt.Errorf("invalid pixel coordinates (%v, %v)", p.X, p.Y)
	}

	bit := uint32(p.X + p.Y*(MAX+1))
	if c.Pixels.Contains(bit) {
		return fmt.Errorf("pixel coordinates (%v, %v) unavailable", p.X, p.Y)
	}

	c.Pixels.Set(bit)
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
		env.Logf("Cleared DApp state")
	case *million.Paint:
		if err := c.Paint(input.Point); err != nil {
			return []byte(FAILURE), err
		}
		env.Logf("Painted pixel %v\n", input.Point)
	default:
		return []byte(FAILURE), fmt.Errorf("invalid input")
	}

	return []byte(SUCCESS), nil
}

func main() {
	eggroll.Roll(&MillionContract{})
}
