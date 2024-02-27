package entities

import (
	"math"
	"project/src/settings"
	"project/utils/mathSys"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Alien struct {
	Entity
	Health          int
	Target_position rl.Vector2
	State           int
}

func NewAlien(x float32, y float32, sprite *rl.Texture2D) Alien {
	alien := Alien{
		Entity: Entity{
			Position: rl.NewVector2(x, y),
			Size:     rl.NewVector2(32, 32),
			Sprite:   sprite,
		},
		Health:          3,
		Target_position: rl.NewVector2(x, y),
		State:           1,
	}

	return alien
}

func (a *Alien) Draw() {
	rl.DrawTextureRec(*a.Sprite, rl.NewRectangle(0, 0, a.Size.X, a.Size.Y), a.Position, rl.White)
}

func (a *Alien) Update() {
	// ---- Lerp ----
	a.Position.X = mathSys.Lerp(a.Position.X, a.Target_position.X, settings.AlienSpeed)
	a.Position.Y = mathSys.Lerp(a.Position.Y, a.Target_position.Y, settings.AlienSpeed)
	// ---- MOVEMENT GRID ----
	if (math.Abs(float64(a.Position.X-a.Target_position.X)) < 1) && (math.Abs(float64(a.Position.Y-a.Target_position.Y)) < 1) {
		a.Position.X = a.Target_position.X
		a.Position.Y = a.Target_position.Y
	}
	//--------------------------------
	if a.Position.X == a.Target_position.X && a.Position.Y == a.Target_position.Y {
		switch a.State {
		case 1:
			if a.Position.X < settings.WINDOW_WIDTH-a.Size.X {
				a.Target_position.X += 32
			} else {
				a.Target_position.Y += 32
				a.State = 2
			}

		case 2:
			if a.Position.Y >= a.Target_position.Y {
				a.Target_position.X -= 32
				a.State = -1
			}
		case -1:
			if a.Position.X > 0 {
				a.Target_position.X -= 32
			} else {
				a.Target_position.Y += 32
				a.State = -2
			}
		case -2:
			if a.Position.Y >= a.Target_position.Y {
				a.Target_position.X += 32
				a.State = 1
			}
		}

	}

}
func (a *Alien) IsCollidingWithShoot(shoot *Shoot) bool {
	return rl.CheckCollisionRecs(
		rl.NewRectangle(a.Position.X, a.Position.Y, a.Size.X, a.Size.Y),
		rl.NewRectangle(shoot.Position.X, shoot.Position.Y, shoot.Size.X, shoot.Size.Y),
	)
}

func (a *Alien) TakeDamage() {
	a.Health--
}
