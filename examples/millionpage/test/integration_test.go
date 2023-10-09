// Copyright (c) Marcel Moura
// SPDX-License-Identifier: MIT (see LICENSE)

package million

import (
	"context"
	"fmt"
	"image"
	"image/color"
	"testing"
	"time"

	"github.com/gligneul/eggroll"
	"github.com/gligneul/eggroll/eggtest"
)

const testTimeout = 300 * time.Second

func TestMillionCtsi(t *testing.T) {
	tester := eggtest.NewIntegrationTester(t)
	defer tester.Close()

	client := eggroll.NewClient()
	lastInputIndex := 0

	hasInputBeenAccepted(t, client, Paint{image.Point{0, 0}, color.RGBA{255, 0, 0, 255}}, lastInputIndex)

	hasInputBeenAccepted(t, client,
		Paint{image.Point{0, 0}, color.RGBA{255, 0, 0, 255}},
		lastInputIndex)

	lastInputIndex += 1
	if _, err := hasInputBeenAccepted(t, client,
		Paint{image.Point{0, 0}, color.RGBA{255, 0, 0, 255}},
		lastInputIndex); err != nil {
		t.Error("Test failed: ", err)
	}

	lastInputIndex += 1
	hasInputBeenAccepted(t, client,
		Paint{image.Point{0, 0}, color.RGBA{255, 0, 0, 255}},
		lastInputIndex)

	lastInputIndex += 1
	if _, err := hasInputBeenAccepted(t, client,
		Paint{image.Point{1000, 0}, color.RGBA{255, 0, 0, 255}},
		lastInputIndex); err != nil {
		t.Error("Test failed: ", err)
	}
}

func hasInputBeenAccepted(
	t *testing.T, client *eggroll.Client,
	input any, lastInputIndex int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	err := client.SendGeneric(ctx, input)
	if err != nil {
		t.Fatalf("failed to send input: %v", err)
	}

	r, err := client.WaitFor(ctx, lastInputIndex)
	if err != nil {
		t.Fatalf("failed to wait for input: %v", err)
	}

	if string(r.Result) == FAILURE {
		return false, fmt.Errorf(r.Logs[0])
	}
	return true, nil
}
