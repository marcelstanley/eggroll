// Copyright (c) Gabriel de Quadros Ligneul
// Copyright (c) Marcel Moura
// SPDX-License-Identifier: MIT (see LICENSE)

// Package that defines the contract logic
package main

import (
	"million"

	"github.com/gligneul/eggroll"
)

func main() {
	eggroll.Roll(&million.Contract{})
}
