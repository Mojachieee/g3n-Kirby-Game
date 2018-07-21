package character

import (
	"time"

	"github.com/mojachieee/g3n-kirby/inputManager"

	"github.com/mojachieee/engine/geometry"
	"github.com/mojachieee/engine/graphic"
	"github.com/mojachieee/engine/material"
	"github.com/mojachieee/engine/math32"
	"github.com/mojachieee/engine/window"
)

// Character is the struct that contains character data
type Character struct {
	model    *geometry.Sphere
	mat      *material.Standard
	Mesh     *graphic.Mesh
	velocity float32
}

// NewCharacter creates a returns a character type
func NewCharacter(model *geometry.Sphere, mat *material.Standard) *Character {
	mat.SetSide(material.SideDouble)
	mat.SetWireframe(true)
	return &Character{
		model,
		mat,
		graphic.NewMesh(model, mat),
		10,
	}
}

// Update is the update function for a character
func (c *Character) Update(deltaTime time.Duration, inputManager *input.Manager) {
	var pos math32.Vector3
	c.checkKeys(inputManager, &pos)

	/* Garbage generation :( */
	vec := c.Mesh.Position()
	vec.Add(pos.MultiplyScalar(c.velocity * float32(deltaTime.Seconds())))
	c.Mesh.SetPositionVec(&vec)
}

// OnKeyDown is the callback for the OnKeyDown event
func (c *Character) checkKeys(i *input.Manager, pos *math32.Vector3) {
	if i.IsKeyDown(window.KeyW) {
		pos.SetZ(-1)
	} else if i.IsKeyDown(window.KeyS) {
		pos.SetZ(1)
	}
	if i.IsKeyDown(window.KeyA) {
		pos.SetX(-1)
	} else if i.IsKeyDown(window.KeyD) {
		pos.SetX(1)
	}
	if i.IsKeyDown(window.KeyE) {
		pos.SetY(1)
	} else if i.IsKeyDown(window.KeyQ) {
		pos.SetY(-1)
	}

}
