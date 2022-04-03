/*
 * @Author: holy
 * @Date: 2022-04-02 01:46:29
 * @LastEditTime: 2022-04-02 09:24:01
 * @LastEditors: holy
 * @Description:
 * @FilePath: \Fyne_learn\03DrawingAndAnimation\_04_Circle\main.go
 */
package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Circle")

	circle := canvas.NewCircle(color.Black)
	circle.StrokeColor = color.Black
	circle.StrokeWidth = 5
	w.SetContent(circle)

	circle1 := canvas.NewCircle(color.Black)
	circle1.StrokeColor = color.Black
	circle1.StrokeWidth = 2
	w.SetContent(circle1)

	w.Resize(fyne.NewSize(100, 100))
	w.ShowAndRun()
}
