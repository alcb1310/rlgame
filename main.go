package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	BG_COLOR   = color.RGBA{207, 255, 240, 0}
	TEXT_COLOR = color.RGBA{105, 105, 255, 255}
	CHAR_COLOR = color.RGBA{255, 161, 84, 255}
)

type Game struct {
	Width    int32
	Height   int32
	FontSize int32

}
func main() {
	fmt.Println("Hello, World!")
	game := &Game{}
	game.Init()

	rl.InitWindow(game.Width, game.Height, "My cool game")
	for !rl.WindowShouldClose() {
		game.Draw()
	}

	rl.CloseWindow()
}

func (g *Game) Init() {
	g.Width = 640
	g.Height = 480
	g.FontSize = 20
}

func (g *Game) Draw() {
	welcomeText := "Congrats! You created your first window!"
	rl.BeginDrawing()
	rl.ClearBackground(BG_COLOR)
	rl.DrawText(welcomeText, g.Width/2-rl.MeasureText(welcomeText, g.FontSize)/2, g.Height/2, g.FontSize, TEXT_COLOR)

	rl.EndDrawing()
}
