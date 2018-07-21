package input

import (
	"github.com/mojachieee/engine/window"
)

// Manager handles key up and down events
type Manager struct {
	keysDown map[window.Key]bool
}

// NewManager returns a new input manager
func NewManager() *Manager {
	return &Manager{keysDown: make(map[window.Key]bool)}
}

// OnKeyUp is the callback function for a key release
func (i *Manager) OnKeyUp(evname string, ev interface{}) {
	if evname == window.OnKeyUp {
		kev := ev.(*window.KeyEvent)
		i.keysDown[kev.Keycode] = false
	}
}

// OnKeyDown is the callback function for a key press
func (i *Manager) OnKeyDown(evname string, ev interface{}) {
	if evname == window.OnKeyDown {
		kev := ev.(*window.KeyEvent)
		i.keysDown[kev.Keycode] = true
	}
}

// IsKeyDown checks if the key is currently being pressed
func (i *Manager) IsKeyDown(keycode window.Key) bool {
	// keycode
	return i.keysDown[keycode]

}
