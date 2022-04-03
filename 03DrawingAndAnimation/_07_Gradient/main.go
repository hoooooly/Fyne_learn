/*
 * @Author: holy
 * @Date: 2022-04-02 10:12:32
 * @LastEditTime: 2022-04-02 10:14:10
 * @LastEditors: holy
 * @Description:
 * @FilePath: \Fyne_learn\03DrawingAndAnimation\_07_Gradient\main.go
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
	w := myApp.NewWindow("Gradient")

	gradient := canvas.NewHorizontalGradient(color.Black, color.Transparent)
	//gradient := canvas.NewRadialGradient(color.White, color.Transparent)
	w.SetContent(gradient)

	w.Resize(fyne.NewSize(100, 100))
	w.ShowAndRun()
}
