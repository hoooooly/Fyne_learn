/*
 * @Author: holy
 * @Date: 2022-04-02 09:26:51
 * @LastEditTime: 2022-04-02 09:30:08
 * @LastEditors: holy
 * @Description:
 * @FilePath: \Fyne_learn\03DrawingAndAnimation\_05_Image\main.go
 */
package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

func main() {
	a := app.New()
	w := a.NewWindow("Image")
	image := canvas.NewImageFromResource(theme.FyneLogo()) // image := canvas.NewImageFromURI(uri)
	// image := canvas.NewImageFromImage(src)
	// image := canvas.NewImageFromReader(reader, name)
	// image := canvas.NewImageFromFile(fileName)
	image.FillMode = canvas.ImageFillOriginal
	w.SetContent(image)

	w.ShowAndRun()
}
