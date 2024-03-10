package game

import (
	"project/src/entities"
	"project/src/settings"
	"project/utils"
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Player        entities.Player
	Asteroids     []*entities.Asteroid
	Shoots        []*entities.Shoot
	Display       rl.RenderTexture2D
	AsteroidTimer utils.Timer
}

func NewGame() Game {
	game := Game{}
	return game
}

func (g *Game) Init() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(settings.WINDOW_WIDTH, settings.WINDOW_HEIGHT, settings.WINDOW_TITLE)
	rl.SetTargetFPS(settings.TARGET_FPS)

	g.Display = rl.LoadRenderTexture(settings.WINDOW_WIDTH, settings.WINDOW_HEIGHT)

}

func (g *Game) Run() {
	//-------- ASSETS --------
	spaceship_asset := rl.LoadTexture("assets/images/nave.png")
	asteroid_asset := rl.LoadTexture("assets/images/asteroid.png")
	shoot_asset := rl.LoadTexture("assets/images/shoot.png")

	g.AsteroidTimer = utils.NewTimer()
	g.AsteroidTimer.Start(2)

	g.Player = entities.NewPlayer(4, float32(settings.WINDOW_HEIGHT/2))
	g.Player.Sprite = &spaceship_asset
	g.Player.ShootSprite = &shoot_asset
	g.Player.Shoots = &g.Shoots
	g.CreateAsteroid(&asteroid_asset)

	for !rl.WindowShouldClose() {
		g.AsteroidTimer.Update()
		g.Player.Update()
		g.UpdateAsteroids()
		g.UpdateShoots()

		if g.AsteroidTimer.IsFinished() {
			g.CreateAsteroid(&asteroid_asset)
			g.AsteroidTimer.Start(float32(rl.GetRandomValue(5, 15)))
		}
		//---- DRAWING ----
		rl.BeginTextureMode(g.Display)
		{
			rl.ClearBackground(rl.Black)

			g.DrawAsteroids()
			g.DrawShoots()
			g.Player.Draw()
		}
		rl.EndTextureMode()
		//---------------------------------------
		rl.BeginDrawing()
		{
			rl.ClearBackground(rl.White)

			//---- Draw the display into the OS window
			rl.DrawTexturePro(
				g.Display.Texture,

				rl.NewRectangle(
					0, 0,
					float32(g.Display.Texture.Width),
					-float32(g.Display.Texture.Height),
				),
				rl.NewRectangle(
					0, 0,
					float32(rl.GetScreenWidth()),
					float32(rl.GetScreenHeight()),
				),

				rl.NewVector2(0, 0),
				0.0,
				rl.White,
			)
		}
		rl.EndDrawing()

	}
}

func (g *Game) Close() {
	rl.CloseWindow()
}

func (g *Game) CreateAsteroid(sprite *rl.Texture2D) {
	asteroid_position := rl.GetRandomValue(0, settings.WINDOW_HEIGHT-96)
	new_asteroid := entities.NewAsteroid(float32(settings.WINDOW_WIDTH-96), float32(asteroid_position), 96/2)
	new_asteroid.Sprite = sprite
	g.Asteroids = append(g.Asteroids, &new_asteroid)
}

func (g *Game) UpdateAsteroids() {
	if len(g.Asteroids) <= 0 {
		return
	}

	for i := 0; i < len(g.Asteroids); i++ {
		(*g.Asteroids[i]).Update()

		if (*g.Asteroids[i]).IsOutsideWindow() {
			g.Asteroids = slices.Delete(g.Asteroids, i, i+1)
		}

	}
}

func (g *Game) DrawAsteroids() {
	if len(g.Asteroids) <= 0 {
		return
	}

	for i := 0; i < len(g.Asteroids); i++ {
		(*g.Asteroids[i]).Draw()
		//(*g.Asteroids[i]).DrawRect()
	}
}

func (g *Game) UpdateShoots() {
	if len(g.Shoots) <= 0 {
		return
	}

	for i := 0; i < len(g.Shoots); i++ {

		if g.Shoots[i] != nil {
			(*g.Shoots[i]).Update()

			g.CheckShootsCollision(i)

			if (*g.Shoots[i]).IsOutsideWindow() {
				g.Shoots = slices.Delete(g.Shoots, i, i+1)

			}
		}
	}
}

func (g *Game) DrawShoots() {
	if len(g.Shoots) <= 0 {
		return
	}

	for i := 0; i < len(g.Shoots); i++ {
		(*g.Shoots[i]).Draw()
	}
}

func (g *Game) CheckShootsCollision(shoot_index int) {
	for i := 0; i < len(g.Asteroids); i++ {
		if len(g.Asteroids) <= 0 {
			return
		}
		println("tooooommmmeeee")
		if (*g.Shoots[shoot_index]).CheckAsteroidCollision(*g.Asteroids[i]) {
			(*g.Shoots[shoot_index]).Position.X += float32(settings.WINDOW_WIDTH) //g.Shoots = append(g.Shoots[:shoot_index], g.Shoots[shoot_index+1:]...) //g.Shoots = slices.Delete(g.Shoots, shoot_index, shoot_index+1)
			(*g.Asteroids[i]).TakeDamage()

			if (*g.Asteroids[i]).Life <= 0 {
				g.Asteroids = slices.Delete(g.Asteroids, i, i+1)
			}

		}
	}
}
