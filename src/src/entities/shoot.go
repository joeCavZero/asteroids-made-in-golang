package entities

import (
	"project/src/settings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Shoot struct {
	Entity
	Direction rl.Vector2
	Speed     float32
}

func NewShoot(x float32, y float32) Shoot {
	return Shoot{
		Entity: Entity{
			Position: rl.NewVector2(x, y),
			Size:     rl.NewVector2(16, 4),
			Rotation: 0.0,
		},
		Speed:     160.0,
		Direction: rl.NewVector2(1, 0),
	}
}

func (s *Shoot) Update() {
	s.Position.X += s.Direction.X * rl.GetFrameTime() * s.Speed
	s.Position.Y += s.Direction.Y * rl.GetFrameTime() * s.Speed
}

func (s *Shoot) Draw() {

	rl.DrawTexturePro(
		*s.Sprite,
		rl.NewRectangle(0, 0, float32(s.Sprite.Width), float32(s.Sprite.Height)),
		rl.NewRectangle(
			s.Position.X+float32(s.Sprite.Width)/2,
			s.Position.Y+float32(s.Sprite.Height)/2,
			float32(s.Sprite.Width),
			float32(s.Sprite.Height),
		),
		rl.NewVector2(float32(s.Sprite.Width)/2, float32(s.Sprite.Height)/2),
		s.Rotation,
		rl.White,
	)
}

func (s *Shoot) IsOutsideWindow() bool {
	return s.Position.X >= float32(settings.WINDOW_WIDTH)
}

func (s *Shoot) CheckAsteroidCollision(a Asteroid) bool {
	return rl.CheckCollisionCircleRec(
		rl.NewVector2(a.Position.X+a.Size.X/2, a.Position.Y+a.Size.Y/2), a.Size.X/2,
		rl.NewRectangle(s.Position.X, s.Position.Y, s.Size.X, s.Size.Y),
	)
}
