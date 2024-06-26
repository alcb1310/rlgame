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

	Character         Character
	CharacterPosition rl.Vector2

	Floor []Floor
}

type Character struct {
	Position rl.Vector2
	Size     rl.Vector2
	Color    rl.Color
	Speed    rl.Vector2
}

type Floor struct {
	Position rl.Vector2
	Size     rl.Vector2
	Color    rl.Color
}

func main() {
	game := &Game{}
	game.Init()

	rl.SetTargetFPS(30)

	rl.InitWindow(game.Width, game.Height, "My cool game")
	for !rl.WindowShouldClose() {
		game.Fall()
		game.Draw()
	}

	rl.CloseWindow()
}

func (g *Game) Init() {
	g.Width = 640
	g.Height = 480
	g.FontSize = 20
	g.Character.Position = rl.Vector2{X: 10, Y: 0}
	g.Character.Size = rl.Vector2{X: 20, Y: 20}
	g.Character.Color = CHAR_COLOR
	g.Character.Speed = rl.Vector2{X: 0, Y: 5}
	g.Floor = []Floor{}
}

func (g *Game) Draw() {
	welcomeText := "Congrats! You created your first window!"
	rl.BeginDrawing()
	rl.ClearBackground(BG_COLOR)
	rl.DrawText(welcomeText, g.Width/2-rl.MeasureText(welcomeText, g.FontSize)/2, g.Height/2, g.FontSize, TEXT_COLOR)

	rl.DrawRectangleV(g.Character.Position, g.Character.Size, g.Character.Color)

	g.Character.Position = rl.Vector2{X: g.Character.Position.X, Y: g.Character.Position.Y}
	fl := Floor{
		Position: rl.Vector2{X: 0, Y: 400},
		Size:     rl.Vector2{X: 100, Y: 10},
		Color:    rl.Color{R: 0, G: 0, B: 0, A: 255},
	}

	fl1 := Floor{
		Position: rl.Vector2{X: 100, Y: 350},
		Size:     rl.Vector2{X: 100, Y: 10},
		Color:    rl.Color{R: 0, G: 0, B: 0, A: 255},
	}
	g.Floor = append(g.Floor, fl)
	g.Floor = append(g.Floor, fl1)

	for _, f := range g.Floor {
		rl.DrawRectangleV(f.Position, f.Size, f.Color)
	}

	// collisions
	for _, floor := range g.Floor {
		if rl.CheckCollisionRecs(
			rl.Rectangle{
				X:      g.Character.Position.X,
				Y:      g.Character.Position.Y,
				Width:  g.Character.Size.X,
				Height: g.Character.Size.Y,
			},
			rl.Rectangle{
				X:      floor.Position.X,
				Y:      floor.Position.Y - floor.Size.Y/2,
				Width:  floor.Size.X,
				Height: floor.Size.Y,
			},
		) {
			g.Character.Position = rl.Vector2{
				X: g.Character.Position.X,
				Y: floor.Position.Y - g.Character.Size.Y - floor.Size.Y/2,
			}
		}
	}

	rl.EndDrawing()
}

func (g *Game) Fall() {
	if int32(g.Character.Position.Y) <= g.Height {
		g.Character.Position = rl.Vector2{X: g.Character.Position.X, Y: g.Character.Position.Y + g.Character.Speed.Y}
		g.Character.Position = rl.Vector2{X: g.Character.Position.X + g.Character.Speed.X, Y: g.Character.Position.Y}
	}
}
