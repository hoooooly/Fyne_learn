/*
 * @Author: holy
 * @Date: 2022-04-04 00:00:51
 * @LastEditTime: 2022-04-04 01:27:41
 * @LastEditors: holy
 * @Description:
 * @FilePath: \Fyne_learn\04ContainersandLayout\_06_Entry\main.go
 */
package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation" //验证器
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Button")

	entry(w)

	w.Resize(fyne.NewSize(500, 500))
	w.ShowAndRun()
}

func entry(w fyne.Window) {
	// 文本输入框
	entry1 := widget.NewEntry()
	entry1.Text = "entry1"
	entry1.TextStyle = fyne.TextStyle{Italic: true}
	entry1.Password = true

	entry2 := widget.NewEntry()
	entry2.TextStyle = fyne.TextStyle{Monospace: true}
	entry2.SetPlaceHolder("placeholder in entry2")
	entry2.MultiLine = true
	entry2.Wrapping = fyne.TextTruncate

	entry3 := widget.NewEntry()
	entry3.TextStyle = fyne.TextStyle{Bold: true} //文本样式
	entry3.Wrapping = fyne.TextWrapBreak          //换行
	entry3.Validator = validation.NewRegexp("^[a-zA-Z]", "Please input Alpha")
	entry3.ActionItem = widget.NewButton("...", func() {
		entry3.SetText("you clicked the actionitem")
	}) //小提示

	entry4 := widget.NewMultiLineEntry()
	entry4.Wrapping = fyne.TextWrapWord
	entry4.Validator = func(s string) error {
		if s != "a" {
			fmt.Println("please input a")
		}
		return nil
	}
	c := container.NewVBox(entry1, entry2, entry3, entry4)
	w.SetContent(c)
}
