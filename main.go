package main

import (
            "fmt"
            "github.com/buger/goterm"
            "math"
            "time"
        )

func main() {
	goterm.Clear()
    
	const pixelAspect = 14.0 / 28.0 // соотношение ширины к высоте символов для терминала mac

	// motion
	step := 0.0
	increasing := true

	for {
        width, height := float64(goterm.Width()), float64(goterm.Height())

        // вычисления координат внутри цикла,
        // чтобы можно было менять размер терминала по ходу выполнения

        // head
        radius3 := 5.0
        centerX3 := width * pixelAspect / 2
        centerY3 := 15.0

        // shaft
        rectX1 := centerX3 - radius3
        rectY1 := centerY3 + 1
        rectX2 := centerX3 + radius3
        rectY2 := rectY1 + 20

        // ball 1
        radius1 := 5.0
        centerX1 := rectX1 + 1
        centerY1 := rectY2

        // ball 2
        radius2 := 6.0
        centerX2 := rectX2
        centerY2 := rectY2

		if increasing {
			step++
		} else {
			step--
		}

		for i := 0.0; i < float64(height); i++ {
			for j := 0.0; j < float64(width); j++ {
				x := j * pixelAspect
				x = x + float64(step)

				if distance(x, centerX1, i, centerY1) <= radius1 {
					fmt.Print("@")
				} else if distance(x, centerX2, i, centerY2) <= radius2 {
					fmt.Print("@")
				} else if is_point_in_area(x, i, rectX1, rectY1, rectX2, rectY2) {
					fmt.Print("#")
				} else if distance(x, centerX3, i, centerY3) <= radius3 {
					fmt.Print("|")
				} else {
					fmt.Print(" ")
				}
			}
		}

		if step == centerX1 - radius1 {
			increasing = false
		} else if step == -centerX1 {
			increasing = true
		}
		time.Sleep(50 * time.Millisecond)
	}
}


func distance(x float64, x0 float64, y float64, y0 float64) float64 {
    return math.Sqrt(math.Pow((x - x0), 2) + math.Pow((y - y0), 2))
}

func is_point_in_area(x float64, y float64, x10 float64, y10 float64, x20 float64, y20 float64) bool {
    return (x10 <= x) && (x <= x20) && (y10 <= y) && (y <= y20)
}