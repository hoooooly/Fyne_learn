/*
 * @Author: holy
 * @Date: 2022-04-02 01:11:42
 * @LastEditTime: 2022-04-02 01:38:40
 * @LastEditors: holy
 * @Description:
 * @FilePath: \Fyne_learn\03DrawingAndAnimation\_01_Rectangle\main.go
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
	w := myApp.NewWindow("Rectangle")

	rect := canvas.NewRectangle(color.Black)
	w.SetContent(rect)

	w.Resize(fyne.NewSize(150, 100))
	w.ShowAndRun()
}
