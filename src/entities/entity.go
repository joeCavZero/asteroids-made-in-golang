package entities

import rl "github.com/gen2brain/raylib-go/raylib"

type Entity struct {
	Position rl.Vector2
	Size     rl.Vector2
	Sprite   *rl.Texture2D
}
