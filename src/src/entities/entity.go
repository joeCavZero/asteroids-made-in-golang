package entities

import rl "github.com/gen2brain/raylib-go/raylib"

type Entity struct {
	Position rl.Vector2
	Size     rl.Vector2
	Rotation float32
	Sprite   *rl.Texture2D
}

func (e *Entity) Draw() {
	rl.DrawTextureEx(*e.Sprite, e.Position, 0.0, 1.0, rl.White)
}

func (e *Entity) DrawRect() {
	rl.DrawLineV(e.Position, rl.NewVector2(e.Position.X+e.Size.X, e.Position.Y), rl.Pink)
	rl.DrawLineV(e.Position, rl.NewVector2(e.Position.X, e.Position.Y+e.Size.Y), rl.Pink)

	rl.DrawLineV(rl.NewVector2(e.Position.X+e.Size.X, e.Position.Y+e.Size.Y), rl.NewVector2(e.Position.X+e.Size.X, e.Position.Y), rl.Pink)
	rl.DrawLineV(rl.NewVector2(e.Position.X+e.Size.X, e.Position.Y+e.Size.Y), rl.NewVector2(e.Position.X, e.Position.Y+e.Size.Y), rl.Pink)
}
