package entities

import rl "github.com/gen2brain/raylib-go/raylib"

type Star struct {
	Entity
}

func NewStar(x float32, y float32) Star {
	return Star{
		Entity: Entity{
			Position: rl.NewVector2(x, y),
			Size:     rl.NewVector2(1, 1),
			Rotation: 0.0,
		},
	}
}

func (s *Star) Update() {
	s.Position.X -= 1
}

func (s *Star) IsOutsideWindow() bool {
	return s.Position.X <= -s.Size.X
}

func (s *Star) Draw() {
	rl.DrawRectangleV(s.Position, s.Size, rl.White)
}
