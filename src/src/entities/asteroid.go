package entities

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Asteroid struct {
	Entity
	Speed float32
	Life  int
}

func NewAsteroid(x float32, y float32, radius float32) Asteroid {
	return Asteroid{
		Entity: Entity{
			Position: rl.NewVector2(x, y),
			Size:     rl.NewVector2(radius*2, radius*2),
			Rotation: 0.0,
		},
		Speed: 50.0,
		Life:  3,
	}
}

func (a *Asteroid) Update() {
	a.Position.X -= a.Speed * rl.GetFrameTime()

}

func (a *Asteroid) Draw() {
	rl.DrawTexturePro(
		*a.Sprite,
		rl.NewRectangle(0, 0, float32(a.Sprite.Width), float32(a.Sprite.Height)),
		rl.NewRectangle(
			a.Position.X+float32(a.Sprite.Width)/2,
			a.Position.Y+float32(a.Sprite.Height)/2,
			float32(a.Sprite.Width),
			float32(a.Sprite.Height),
		),
		rl.NewVector2(float32(a.Sprite.Width)/2, float32(a.Sprite.Height)/2),
		a.Rotation,
		rl.White,
	)
}

func (a *Asteroid) IsOutsideWindow() bool {
	return a.Position.X <= -a.Size.X
}

func (a *Asteroid) TakeDamage() {
	a.Life -= 1
}
