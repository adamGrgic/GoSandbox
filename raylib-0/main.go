package main

import(
    gui "github.com/gen2brain/raylib-go/raygui"
    rl "github.com/gen2brain/raylib-go/raylib"
    "fmt"
)

func main() {
	rl.InitWindow(1200, 750, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

    var button bool
    var button2 bool

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

        rl.DrawRectangle(int32(250), int32(350), int32(100), int32(40))

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 100, 200, 20, rl.LightGray)

        button2 = gui.Button(rl.NewRectangle(50, 150, 100, 40), "Click 2")
        if button2 {
            fmt.Println("Clicked on button 2")
        }

        button = gui.Button(rl.NewRectangle(-500, 150, 100, 40), "Click")
		if button {
			fmt.Println("Clicked on button")
		}
		rl.EndDrawing()
	}
}
