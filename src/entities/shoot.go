package entities

import rl "github.com/gen2brain/raylib-go/raylib"

type Shoot struct {
	Entity
	Speed float32
	Angle float32
}

func NewShoot(x float32, y float32, sprite *rl.Texture2D) Shoot {
	shoot := Shoot{
		Entity: Entity{
			Position: rl.NewVector2(x, y),
			Size:     rl.NewVector2(4, 16),
			Sprite:   sprite,
		},
		Speed: 200,
		Angle: 0.0,
	}

	return shoot
}
func (s *Shoot) IsOutsideWindow() bool {
	return s.Position.Y < -s.Size.Y
}
func (s *Shoot) Update() {
	s.Position.Y += -1 * s.Speed * rl.GetFrameTime()
}

func (s *Shoot) Draw() {
	rl.DrawTexturePro(
		*s.Sprite,
		rl.NewRectangle(0, 0, s.Size.X, s.Size.Y),
		rl.NewRectangle(s.Position.X, s.Position.Y, s.Size.X, s.Size.Y),
		rl.NewVector2(0, 0),
		s.Angle,
		rl.White,
	)
}
