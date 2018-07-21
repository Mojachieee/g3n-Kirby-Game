// Copyright 2016 The G3N Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This is a minimum G3N application showing how to create a window,
// a scene, add some 3D objects to the scene and render it.
// For more complete demos please see: https://github.com/g3n/g3nd
package main

import (
	"math"

	"github.com/mojachieee/engine/geometry"
	"github.com/mojachieee/engine/light"
	"github.com/mojachieee/engine/material"
	"github.com/mojachieee/engine/math32"
	"github.com/mojachieee/engine/util/application"
	"github.com/mojachieee/engine/window"
	"github.com/mojachieee/g3n-kirby/character"
	input "github.com/mojachieee/g3n-kirby/inputManager"
)

type game struct {
	chars        []*character.Character
	app          *application.Application
	inputManager *input.Manager
}

func (g game) Update(evname string, ev interface{}) {
	for _, elm := range g.chars {
		elm.Update(g.app.FrameDelta(), g.inputManager)
	}
}

func main() {

	app, err := application.Create(application.Options{
		Title:      "Kirby Game",
		Width:      800,
		Height:     600,
		Fullscreen: false,
		TargetFPS:  60,
	})
	if err != nil {
		panic(err)
	}

	char := character.NewCharacter(
		geometry.NewSphere(2, 16, 16, 0, math.Pi*2, 0, math.Pi),
		material.NewStandard(math32.NewColor("Pink")),
	)

	g := game{
		app:          app,
		chars:        append([]*character.Character{}, char),
		inputManager: input.NewManager(),
	}

	g.app.Scene().Add(light.NewAmbient(&math32.Color{1.0, 1.0, 1.0}, 0.5))

	g.app.Scene().Add(char.Mesh)
	g.app.Subscribe(application.OnBeforeRender, g.Update)
	g.app.Window().Subscribe(window.OnKeyDown, g.inputManager.OnKeyDown)
	g.app.Window().Subscribe(window.OnKeyUp, g.inputManager.OnKeyUp)
	// g.app.Window().Subscribe(window.OnKeyRepeat, char.OnInput)
	err = g.app.Run()
	if err != nil {
		panic(err)
	}
}
