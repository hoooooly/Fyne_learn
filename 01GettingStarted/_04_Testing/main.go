/*
 * @Author: holy
 * @Date: 2022-04-01 01:29:22
 * @LastEditTime: 2022-04-01 01:38:07
 * @LastEditors: holy
 * @Description:
 * @FilePath: \Fyne_learn\01GettingStarted\_04_Testing\main.go
 */
package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func makeUI() (*widget.Label, *widget.Entry) {
	return widget.NewLabel("Hello world"), widget.NewEntry()
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello Person")

	w.ShowAndRun()
}
