/*
 * @Author: holy
 * @Date: 2022-04-01 02:02:30
 * @LastEditTime: 2022-04-01 02:08:52
 * @LastEditors: holy
 * @Description:容器和布局
 * @FilePath: \Fyne_learn\02ExploringFyne\_02_ContainerandLayouts\main.go
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
	w := a.NewWindow("Container")
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}

	text1 := canvas.NewText("Hello", green)
	text2 := canvas.NewText("World", green)
	//content := container.NewWithoutLayout(text1, text2)
	content := container.New(layout.NewGridLayout(2), text1, text2)

	w.SetContent(content)
	w.ShowAndRun()
}
