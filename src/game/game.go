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
	Alien_slice []entities.Alien
	Shoot_slice []entities.Shoot
	Spawn_side  int
}

func NewGame() Game {
	game := Game{}
	game.Spawn_side = rand.Intn(2)
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
	alienSprite := rl.LoadTexture("assets/alien.png")
	playerSprite := rl.LoadTexture("assets/player.png")
	shootSprite := rl.LoadTexture("assets/shoot.png")
	naveMaeSprite := rl.LoadTexture("assets/nave-mae.png")
	display := rl.LoadRenderTexture(settings.WINDOW_WIDTH, settings.WINDOW_HEIGHT)

	// ---- ALIEN CREATION ----

	new_alien := entities.NewAlien(-32.0, 0.0, &alienSprite)

	if g.Spawn_side == 0 {
		new_alien.State = 1
		new_alien.Position = rl.NewVector2(-32, 0)
		new_alien.Target_position.X = 0
	} else if g.Spawn_side == 1 {
		new_alien.State = -1
		new_alien.Position = rl.NewVector2(settings.WINDOW_WIDTH, 0)
		new_alien.Target_position.X = settings.WINDOW_WIDTH - 32
	}
	g.Alien_slice = append(g.Alien_slice, new_alien)

	player := entities.NewPlayer(settings.WINDOW_WIDTH/2, settings.WINDOW_HEIGHT-32, &playerSprite)
	player.Shoot_sprite = &shootSprite
	player.Shoots_slice = &g.Shoot_slice

	timer_ := timer.NewTimer()
	// ---- MAIN LOOP ----
	for !rl.WindowShouldClose() {

		for i := 0; i < len(g.Alien_slice); i++ {
			g.Alien_slice[i].Update()
		}
		for i := 0; i < len(g.Shoot_slice); i++ {
			g.Shoot_slice[i].Update()

			if g.Shoot_slice[i].IsOutsideWindow() {
				g.Shoot_slice = slices.Delete(g.Shoot_slice, i, i+1)

			}
		}
		player.Update()
		// ============= DRAWING ON DISPLAY =============
		rl.BeginTextureMode(display)
		rl.ClearBackground(rl.Black)
		//DrawGrid(rl.Gray)
		rl.DrawTexture(naveMaeSprite, 0, settings.WINDOW_HEIGHT-64, rl.White)

		player.Draw()

		for i := 0; i < len(g.Alien_slice); i++ {
			g.Alien_slice[i].Draw()
		}
		for i := 0; i < len(g.Shoot_slice); i++ {
			g.Shoot_slice[i].Draw()
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
