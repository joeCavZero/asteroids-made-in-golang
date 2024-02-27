package game

import (
	"math/rand"
	"project/src/entities"
	"project/src/settings"
	"project/utils/timer"
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	AlienSlice  []entities.Alien
	ShootsSlice []entities.Shoot
	SpawnSide   int
}

func NewGame() Game {
	game := Game{}
	game.SpawnSide = rand.Intn(2)
	return game
}

func (g *Game) Init() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(settings.WINDOW_WIDTH, settings.WINDOW_HEIGHT, "Space Invaders")
	rl.SetTargetFPS(60)
	rl.SetExitKey(rl.KeyNull)
}

func (g *Game) GameLoop() {
	//======== ASSETS ========
	alienSprite := rl.LoadTexture("assets/images/alien.png")
	playerSprite := rl.LoadTexture("assets/images/player.png")
	shootSprite := rl.LoadTexture("assets/images/shoot.png")
	naveMaeSprite := rl.LoadTexture("assets/images/nave-mae.png")
	display := rl.LoadRenderTexture(settings.WINDOW_WIDTH, settings.WINDOW_HEIGHT)

	player := entities.NewPlayer(settings.WINDOW_WIDTH/2, settings.WINDOW_HEIGHT-32, &playerSprite)
	player.ShootSprite = &shootSprite
	player.ShootsSlice = &g.ShootsSlice

	alienTimer := timer.NewTimer()

	// ---- MAIN LOOP ----
	for !rl.WindowShouldClose() {
		// ============= UPDATE =============
		if alienTimer.IsFinished() {
			newAlien := entities.NewAlien(-32.0, 0.0, &alienSprite)
			g.AlienSlice = slices.Insert(g.AlienSlice, len(g.AlienSlice), newAlien)
			alienTimer.Start(3.0)
		}

		for i := 0; i < len(g.AlienSlice); i++ {
			g.AlienSlice[i].Update()
			if g.AlienSlice[i].Health <= 0 {

				g.AlienSlice = slices.Delete(g.AlienSlice, i, i+1)
			}
			if len(g.AlienSlice) > 0 {
				for j := 0; j < len(g.ShootsSlice); j++ {

					if g.AlienSlice[i].IsCollidingWithShoot(&g.ShootsSlice[j]) {
						g.AlienSlice[i].TakeDamage()

						g.ShootsSlice = slices.Delete(g.ShootsSlice, j, j+1)
					}
				}
			}
		}
		for i := 0; i < len(g.ShootsSlice); i++ {
			g.ShootsSlice[i].Update()

			if g.ShootsSlice[i].IsOutsideWindow() {
				g.ShootsSlice = slices.Delete(g.ShootsSlice, i, i+1)

			}
		}
		alienTimer.Update()
		player.Update()

		// ============= DRAWING ON DISPLAY =============
		rl.BeginTextureMode(display)
		rl.ClearBackground(rl.Black)
		//DrawGrid(rl.Gray)
		rl.DrawTexture(naveMaeSprite, 0, settings.WINDOW_HEIGHT-64, rl.White)

		player.Draw()

		for i := 0; i < len(g.AlienSlice); i++ {
			g.AlienSlice[i].Draw()
		}
		for i := 0; i < len(g.ShootsSlice); i++ {
			g.ShootsSlice[i].Draw()
		}

		rl.EndTextureMode()

		// ---- DRAWING DISPLAY ON WINDOW ----
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawTexturePro(display.Texture,
			rl.NewRectangle(0, 0, float32(display.Texture.Width), -float32(display.Texture.Height)),
			rl.NewRectangle(0, 0, float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight())),
			rl.NewVector2(0, 0),
			0,
			rl.White,
		)

		rl.EndDrawing()
	}

	rl.UnloadTexture(alienSprite)
	rl.UnloadTexture(playerSprite)
	rl.UnloadTexture(shootSprite)
	rl.UnloadTexture(naveMaeSprite)

}

func DrawGrid(color rl.Color) {
	for i := 0; i < settings.WINDOW_WIDTH/32; i++ {
		rl.DrawLine(int32(i*32), 0, int32(i*32), settings.WINDOW_HEIGHT, color)
	}
	for i := 0; i < settings.WINDOW_HEIGHT/32; i++ {
		rl.DrawLine(0, int32(i*32), settings.WINDOW_WIDTH, int32(i*32), color)
	}
}

func (g *Game) Close() {
	rl.CloseWindow()
}
