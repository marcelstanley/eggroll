// Copyright (c) Gabriel de Quadros Ligneul
// Copyright (c) Marcel Moura
// SPDX-License-Identifier: MIT (see LICENSE)

// Package tha defines the client logic
package main

import (
	"context"
	"image"
	"image/color"
	"log"
	"million"

	"github.com/gligneul/eggroll"
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Must[T any](obj T, err error) T {
	Check(err)
	return obj
}

// TODO Create a backing image for the DApp state and save it to a file
// Map the contract addresses to it to provide the full picture
type Address struct {
	Value [20]byte
}

func main() {
	ctx := context.Background()
	client := eggroll.NewClient()

	inputs := []any{
		million.Init{},
		million.Paint{image.Point{0, 0}, color.RGBA{255, 0, 0, 255}},
		million.Paint{image.Point{999, 999}, color.RGBA{255, 255, 0, 0}},
	}
	for _, input := range inputs {
		log.Printf("Sending input %#v\n", input)
		Check(client.SendGeneric(ctx, input))
	}

	log.Println("Waiting for inputs to be processed")
	Check(client.WaitFor(ctx, 2))

	Check(client.Sync(ctx))
	var contract million.Contract
	Check(client.ReadState(&contract))

	log.Println("Map: ", contract.Pixels.ToBytes())

	logs := Must(client.Logs(ctx))
	log.Println("Logs:")
	for _, msg := range logs {
		log.Print(">", msg)
	}
}
