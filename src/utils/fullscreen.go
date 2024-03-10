package utils

import (
	"project/src/settings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func SwitchFullscreen() {
	if rl.IsWindowFullscreen() {
		rl.ToggleFullscreen()
		rl.SetWindowSize(int(settings.WINDOW_WIDTH*2), int(settings.WINDOW_HEIGHT*2))
		rl.SetWindowPosition(int(settings.WINDOW_WIDTH/4), int(settings.WINDOW_HEIGHT/4))
	} else {
		monitor := rl.GetCurrentMonitor()

		rl.SetWindowSize(rl.GetMonitorWidth(monitor), rl.GetMonitorWidth(monitor))
		rl.ToggleFullscreen()
	}
}
