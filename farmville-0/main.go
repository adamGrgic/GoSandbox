package main

import (
	"fmt"
	"os"
    "github.com/gdamore/tcell/v2"
    "log"
)

// main event loop
//func step() {
    //fmt.Println("DEBUG: step")


//}
func drawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Fill background
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			s.SetContent(col, row, ' ', nil, style)
		}
	}

	// Draw borders
	for col := x1; col <= x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}

	// Only draw corners if necessary
	if y1 != y2 && x1 != x2 {
		s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
		s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
		s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
		s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
	}

	drawText(s, x1+1, y1+1, x2-1, y2-1, style, text)
}

func main() {
    fmt.Println("Starting up Farmville module...")

    var name string

    fmt.Print("Enter your name: ")
    _, err := fmt.Scan(&name)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
    }

    // debug_lines := fmt.Sprintf("DEBUG: lines: %d", lines)
    // fmt.Println(debug_lines)

    fmt.Println("Hello,", name)

    fmt.Println("Starting new game... ")

    s, err := tcell.NewScreen()
    if err != nil {
        log.Fatalf("%+v", err)
    }
    if err := s.Init(); err != nil {
        log.Fatalf("%+v", err)
    }

    boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)
    // Set default text style
    defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
    s.SetStyle(defStyle)

    // Clear screen
    s.Clear()

    // Draw initial boxes
	drawBox(s, 1, 1, 42, 7, boxStyle, "Click and drag to draw a box")
	drawBox(s, 5, 9, 32, 14, boxStyle, "Press C to reset")
//    gameRunning := true

 //   for (gameRunning) {
 //       // fmt.Println("DEBUG: Game is running")
 //       step()
 //   }
}
