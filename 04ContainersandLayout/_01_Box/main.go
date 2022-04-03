/*
 * @Author: holy
 * @Date: 2022-04-02 10:43:46
 * @LastEditTime: 2022-04-02 23:41:16
 * @LastEditors: holy
 * @Description:
 * @FilePath: \Fyne_learn\04ContainersandLayout\_01_Box\main.go
 */
package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	a := app.New()
	w := a.NewWindow("Box layout")

	text1 := canvas.NewText("hello", color.Black)
	text2 := canvas.NewText("There", color.Black)
	text3 := canvas.NewText("right", color.Black)
	content := container.New(layout.NewHBoxLayout(), text1, text2, layout.NewSpacer(), text3)

	text4 := canvas.NewText("centered", color.Black)
	centered := container.New(layout.NewHBoxLayout(), text4, layout.NewSpacer(), text3)
	w.SetContent(container.New(layout.NewVBoxLayout(), content, centered))
	w.ShowAndRun()
}
