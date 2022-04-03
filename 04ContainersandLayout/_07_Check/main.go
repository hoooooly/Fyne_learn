/*
 * @Author: holy
 * @Date: 2022-04-04 01:31:32
 * @LastEditTime: 2022-04-04 02:00:58
 * @LastEditors: holy
 * @Description:
 * @FilePath: \Fyne_learn\04ContainersandLayout\_07_Check\main.go
 */
package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Button")

	check(w)
	w.Resize(fyne.NewSize(500, 500))
	w.ShowAndRun()
}

func check(w fyne.Window) {
	lblmsg := widget.NewLabel("")

	chk1 := widget.NewCheck("Sun", nil)
	chk1.OnChanged = func(b bool) {
		if b && !strings.Contains(lblmsg.Text, chk1.Text) {
			lblmsg.SetText(lblmsg.Text + "," + chk1.Text)
		}
		if strings.Contains(lblmsg.Text, chk1.Text) && !b {
			lblmsg.SetText(strings.Replace(lblmsg.Text, ","+chk1.Text, "", -1))
		}
	}

	chk2 := widget.NewCheck("Moon", nil)
	chk2.OnChanged = func(b bool) {
		if b && !strings.Contains(lblmsg.Text, chk2.Text) {
			lblmsg.SetText(lblmsg.Text + "," + chk2.Text)
		}
		if strings.Contains(lblmsg.Text, chk2.Text) && !b {
			lblmsg.SetText(strings.Replace(lblmsg.Text, ","+chk2.Text, ",", -1))
		}
	}
	c := container.NewHBox(lblmsg, chk1, chk2)
	w.SetContent(c)
}
