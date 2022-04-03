/*
 * @Author: holy
 * @Date: 2022-04-04 00:00:51
 * @LastEditTime: 2022-04-04 01:02:49
 * @LastEditors: holy
 * @Description:
 * @FilePath: \Fyne_learn\04ContainersandLayout\_05_Button\main.go
 */
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Button")

	btnDemo(w)

	w.Resize(fyne.NewSize(500, 500))
	w.ShowAndRun()
}

func btnDemo(w fyne.Window) {
	lblmsg := widget.NewLabel("")
	btn1 := widget.NewButton("btn1", func() {
		lblmsg.SetText("Btn1 clicked")
	})
	btn1.Importance = widget.MediumImportance
	btn2 := widget.NewButtonWithIcon("Btn2", theme.ComputerIcon(), func() {
		lblmsg.SetText("Btn2 clicked")
	})
	btn2.Alignment = widget.ButtonAlignCenter
	btn2.Importance = widget.HighImportance
	btn3 := widget.NewButtonWithIcon("Btn3", theme.ConfirmIcon(), func() {
		lblmsg.SetText("Btn3 clicked")
	})
	//自定义button
	btn4 := NewMyButton("myBtn4", func() {
		lblmsg.SetText("myBtn4 clicked")
	})

	btn5 := NewMyButton("myBtn5", func() {
		lblmsg.SetText("myBtn5 clicked, btn4 Hidden, click again btn4 appear!")
		btn4.Button.Hidden = !btn4.Button.Hidden
	})

	// 设置容器，放置控件
	c := container.NewVBox(lblmsg, btn1, btn2, btn3, btn4, btn5)
	w.SetContent(c)
}

//自定义button
type MyButton struct {
	widget.Button
	CCursor desktop.Cursor
}

func NewMyButton(label string, tapped func()) *MyButton {
	button := &MyButton{}
	button.Button.Text = label
	button.Button.OnTapped = tapped

	return button
}
