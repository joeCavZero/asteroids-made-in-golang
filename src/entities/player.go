package entities

import (
	"project/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Entity
	Speed     float32
	Direction rl.Vector2

	ShootSprite        *rl.Texture2D
	Shoots             *[]*Shoot
	ShootSide          int
	ShootTimerDebounce utils.Timer
	ShootDebounceTime  float32
}

func NewPlayer(x float32, y float32) Player {
	return Player{
		Entity: Entity{
			Position: rl.NewVector2(x, y),
			Size:     rl.NewVector2(32, 32),
			Rotation: 0.0,
		},
		Speed:              100.0,
		Direction:          rl.NewVector2(0.0, 0.0),
		ShootTimerDebounce: utils.NewTimer(),
		ShootDebounceTime:  0.5,
		ShootSide:          1,
	}
}

func (p *Player) Draw() {
	//rl.DrawTextureEx(*p.Sprite, p.Position, 0.0, 1.0, rl.White)
	p.Rotation = rl.Lerp(p.Rotation, p.Direction.Y*5, 0.1)
	rl.DrawTexturePro(
		*p.Sprite,
		rl.NewRectangle(0, 0, float32(p.Sprite.Width), float32(p.Sprite.Height)),
		rl.NewRectangle(
			p.Position.X+float32(p.Sprite.Width)/2,
			p.Position.Y+float32(p.Sprite.Height)/2,
			float32(p.Sprite.Width),
			float32(p.Sprite.Height),
		),
		rl.NewVector2(float32(p.Sprite.Width)/2, float32(p.Sprite.Height)/2),
		p.Rotation,
		rl.White,
	)
}

func (p *Player) Update() {
	p.input()
	p.move()
	p.ShootTimerDebounce.Update()
}

func (p *Player) input() {
	if rl.IsKeyDown(rl.KeyW) {
		p.Direction.Y = -1
	} else if rl.IsKeyDown(rl.KeyS) {
		p.Direction.Y = 1
	} else {
		p.Direction.Y = 0
	}

	if rl.IsKeyDown(rl.KeyA) {
		p.Direction.X = -1
	} else if rl.IsKeyDown(rl.KeyD) {
		p.Direction.X = 1
	} else {
		p.Direction.X = 0
	}

	p.Direction = rl.Vector2Normalize(p.Direction)

	if rl.IsKeyDown(rl.KeySpace) && p.ShootTimerDebounce.IsFinished() {
		p.CreateShoot()
		p.ShootTimerDebounce.Start(p.ShootDebounceTime)

	}
}

func (p *Player) move() {
	p.Position.X += p.Direction.X
	p.Position.Y += p.Direction.Y
}

func (p *Player) CreateShoot() {
	newShoot := NewShoot(p.Position.X+p.Size.X/2, p.Position.Y+p.Size.Y/2)
	newShoot.Position.Y += float32(p.ShootSide)*8 - newShoot.Size.Y/2
	p.ShootSide *= -1

	newShoot.Sprite = p.ShootSprite
	*p.Shoots = append(*p.Shoots, &newShoot)
}
