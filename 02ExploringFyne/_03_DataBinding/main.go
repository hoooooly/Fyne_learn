/*
 * @Author: holy
 * @Date: 2022-04-02 00:59:47
 * @LastEditTime: 2022-04-02 01:08:47
 * @LastEditors: holy
 * @Description:
 * @FilePath: \Fyne_learn\02ExploringFyne\_03_DataBinding\main.go
 */
package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("hello")
	w.Resize(fyne.NewSize(100, 200))

	str := binding.NewString()
	go func() {
		dots := "....."
		for i := 10; i >= 0; i-- {
			for i := 1; i <= 3; i++ {
				str.Set("Count down" + dots[:i])
				time.Sleep(time.Second)
			}
		}

		str.Set("Blast off")
	}()
	w.SetContent(widget.NewLabelWithData(str))
	w.ShowAndRun()
}
