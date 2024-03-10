package game

import (
	"math/rand"
	"project/src/entities"
	"project/src/settings"
	"project/utils"
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Player    entities.Player
	Asteroids []*entities.Asteroid
	Shoots    []*entities.Shoot
	Stars     []entities.Star
	Display   rl.RenderTexture2D

	AsteroidTimer utils.Timer
	StarTimer     utils.Timer

	PlayerLife int
}

func NewGame() Game {
	game := Game{PlayerLife: 5}
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
	heart_asset := rl.LoadTexture("assets/images/heart.png")

	g.AsteroidTimer = utils.NewTimer()
	g.AsteroidTimer.Start(2)

	g.StarTimer = utils.NewTimer()
	g.StarTimer.Start(0.1)

	g.Player = entities.NewPlayer(4, float32(settings.WINDOW_HEIGHT/2))
	{
		g.Player.Sprite = &spaceship_asset
		g.Player.ShootSprite = &shoot_asset
		g.Player.Shoots = &g.Shoots
		g.Player.Life = &g.PlayerLife
	}
	g.CreateAsteroid(&asteroid_asset)
	g.CreateFirstStars()
	for !rl.WindowShouldClose() {
		g.AsteroidTimer.Update()
		g.StarTimer.Update()

		g.Player.Update()
		g.UpdateAsteroids()
		g.UpdateShoots()
		g.UpdateStars()

		if g.AsteroidTimer.IsFinished() {
			g.CreateAsteroid(&asteroid_asset)
			g.AsteroidTimer.Start(float32(rl.GetRandomValue(3, 6)))
		}

		if g.StarTimer.IsFinished() {
			g.CreateStar()
			g.StarTimer.Start(float32(rl.GetRandomValue(1, 3) / 3))
		}

		if g.PlayerLife <= 0 {
			g.Restart()
		}
		//---- DRAWING ----
		rl.BeginTextureMode(g.Display)
		{
			rl.ClearBackground(rl.Black)
			g.DrawStars()
			g.DrawAsteroids()
			g.DrawShoots()

			g.Player.Draw()
			g.DrawHearts(heart_asset)
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
	new_asteroid := entities.NewAsteroid(float32(settings.WINDOW_WIDTH+16), float32(asteroid_position), ([]float32{16, 32, 48})[rand.Intn(3)])
	new_asteroid.Position.Y = float32(rl.GetRandomValue(0, settings.WINDOW_HEIGHT-int32(new_asteroid.Size.Y)))
	new_asteroid.Sprite = sprite
	g.Asteroids = append(g.Asteroids, &new_asteroid)
}

func (g *Game) UpdateAsteroids() {
	if len(g.Asteroids) <= 0 {
		return
	}

	for i := 0; i < len(g.Asteroids); i++ {
		(*g.Asteroids[i]).Update()
		g.IsPlayerCollidingAsteroid(i)

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

		if (*g.Shoots[shoot_index]).CheckAsteroidCollision(*g.Asteroids[i]) {
			(*g.Shoots[shoot_index]).Position.X += float32(settings.WINDOW_WIDTH) //g.Shoots = append(g.Shoots[:shoot_index], g.Shoots[shoot_index+1:]...) //g.Shoots = slices.Delete(g.Shoots, shoot_index, shoot_index+1)
			(*g.Asteroids[i]).TakeDamage(1)

			if (*g.Asteroids[i]).Life <= 0 {
				g.Asteroids = slices.Delete(g.Asteroids, i, i+1)
			}

		}
	}
}

func (g *Game) DrawHearts(heartTexture rl.Texture2D) {
	for i := 0; i < g.PlayerLife; i++ {
		//rl.DrawTextureEx(heartTexture, rl.NewVector2(float32(i*24), 0), 0.0, 1.0, rl.White)
		rl.DrawTexturePro(heartTexture,
			rl.NewRectangle(0, 0, 32, 32),
			rl.NewRectangle(float32(i*32), 4, 32, 32),
			rl.NewVector2(0, 0), 0.0, rl.White,
		)
	}
}

func (g *Game) IsPlayerCollidingAsteroid(aster_index int) {
	actual_asteroid := *g.Asteroids[aster_index]
	if rl.CheckCollisionCircleRec(
		rl.NewVector2(actual_asteroid.Position.X+actual_asteroid.Size.X/2, actual_asteroid.Position.Y+actual_asteroid.Size.Y/2),
		actual_asteroid.Size.X/2,
		rl.NewRectangle(g.Player.Position.X, g.Player.Position.Y, g.Player.Size.X, g.Player.Size.Y),
	) {
		g.PlayerLife -= 1
		(*g.Asteroids[aster_index]).Position.X -= 1000
	}
}

func (g *Game) Restart() {
	g.Asteroids = nil
	g.Player.Position = rl.NewVector2(4, float32(settings.WINDOW_HEIGHT/2))
	g.PlayerLife = 5
}
func (g *Game) CreateFirstStars() {
	for i := 0; i < 100; i++ {
		g.CreateStar()
		g.Stars[i].Position.X = float32(rl.GetRandomValue(0, settings.WINDOW_WIDTH))
	}

}
func (g *Game) CreateStar() {
	star_position := rl.GetRandomValue(0, settings.WINDOW_HEIGHT)
	new_star := entities.NewStar(float32(settings.WINDOW_WIDTH+16), float32(star_position))
	g.Stars = append(g.Stars, new_star)
}

func (g *Game) UpdateStars() {
	if len(g.Stars) <= 0 {
		return
	}

	for i := 0; i < len(g.Stars); i++ {
		g.Stars[i].Update()

		if g.Stars[i].IsOutsideWindow() {
			g.Stars = slices.Delete(g.Stars, i, i+1)
		}
	}
}

func (g *Game) DrawStars() {
	if len(g.Stars) <= 0 {
		return
	}

	for i := 0; i < len(g.Stars); i++ {
		g.Stars[i].Draw()
	}
}
