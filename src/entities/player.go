package entities

import (
	"math"
	"project/src/settings"
	"project/utils/timer"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Entity
	Health        int
	Motion        rl.Vector2
	Speed         float32
	ShootsSlice   *[]Shoot
	ShootSprite   *rl.Texture2D
	NextShootSide int
	ShootDebounce timer.Timer
}

func NewPlayer(x float32, y float32, sprite *rl.Texture2D) Player {
	player := Player{
		Entity: Entity{
			Position: rl.NewVector2(x, y),
			Size:     rl.NewVector2(32, 32),
			Sprite:   sprite,
		},
		Health:        5,
		Motion:        rl.NewVector2(0, 0),
		Speed:         105,
		NextShootSide: 0,
		ShootDebounce: timer.NewTimer(),
	}

	return player
}

func (p *Player) Input() {
	if rl.IsKeyDown(rl.KeyW) {
		p.Motion.Y = -1
	} else if rl.IsKeyDown(rl.KeyS) {
		p.Motion.Y = 1
	} else {
		p.Motion.Y = 0
	}

	if rl.IsKeyDown(rl.KeyA) {
		p.Motion.X = -1
	} else if rl.IsKeyDown(rl.KeyD) {
		p.Motion.X = 1
	} else {
		p.Motion.X = 0
	}

	p.Motion.X = rl.Clamp(p.Motion.X, -1, 1)
	p.Motion.Y = rl.Clamp(p.Motion.Y, -1, 1)
	p.Motion = rl.Vector2Normalize(p.Motion)

	if rl.IsKeyDown(rl.KeyQ) && p.ShootDebounce.IsFinished() {
		p.Shoot(p.Position.X, p.Position.Y)
		p.ShootDebounce.Start(0.3)
	}
}
func (p *Player) Move() {
	p.Position.X += p.Motion.X * p.Speed * rl.GetFrameTime()
	p.Position.Y += p.Motion.Y * p.Speed * rl.GetFrameTime()

	p.Position.X = float32(math.Max(0.0, math.Min(float64(p.Position.X), float64(settings.WINDOW_WIDTH)-float64(p.Size.X))))
	p.Position.Y = float32(math.Max(0.0, math.Min(float64(p.Position.Y), float64(settings.WINDOW_HEIGHT)-float64(p.Size.Y))))
}
func (p *Player) Update() {
	p.Input()
	p.Move()

	p.ShootDebounce.Update()
}

func (p *Player) Draw() {

	rl.DrawTexturePro(*p.Sprite,
		rl.NewRectangle(0, 0, p.Size.X, p.Size.Y),
		rl.NewRectangle(p.Position.X, p.Position.Y, p.Size.X, p.Size.Y),
		rl.NewVector2(0, 0),
		0.0, rl.White,
	)
}

//=========================================

func (p *Player) Shoot(x float32, y float32) {
	var newShoot Shoot

	switch p.NextShootSide {
	case 0:
		newShoot = NewShoot(p.Position.X+2.0, p.Position.Y, p.ShootSprite)
		p.NextShootSide = 1

	case 1:
		newShoot = NewShoot(p.Position.X+p.Size.X-6.0, p.Position.Y, p.ShootSprite)
		p.NextShootSide = 0
	}

	*p.ShootsSlice = append(*p.ShootsSlice, newShoot)

}
