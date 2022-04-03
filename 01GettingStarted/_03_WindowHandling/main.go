/*
 * @Author: holy
 * @Date: 2022-04-01 01:17:11
 * @LastEditTime: 2022-04-01 01:26:57
 * @LastEditors: holy
 * @Description:
 * @FilePath: \Fyne_learn\01GettingStarted\_03_WindowHandling\main.go
 */
package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func updateTime2Label(l *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05") //设置时间
	l.SetText(formatted)                             //设置标签文本
}

func main() {
	myApp := app.New()                   //创建应用程序
	myWindow := myApp.NewWindow("Hello") //创建窗口
	labelClock := widget.NewLabel("")    //设置Label控件
	updateTime2Label(labelClock)
	myWindow.SetContent(widget.NewLabel("Hello")) //在窗口上设置控件
	myWindow.SetContent(labelClock)
	go func() {
		for range time.Tick(time.Second) {
			updateTime2Label(labelClock)
		}
	}()

	w2 := myApp.NewWindow("Larger Window") //另一个窗口
	w2.SetContent(widget.NewLabel("Lager window"))
	w2.Resize(fyne.NewSize(400, 500)) //设置宽高400x500
	w2.Show()

	w2.SetContent(widget.NewButton("Open New", func() {
		w3 := myApp.NewWindow("Third Window")
		w3.SetContent(widget.NewLabel("Third content"))
		w3.Resize(fyne.NewSize(200, 300))
		w3.Show()
	})) //添加按钮，绑定一个点击事件，点击后创建一个新的窗口
	// myWindow.Show() //显示窗口
	// myApp.Run()     //运行应用程序
	myWindow.ShowAndRun() //显示窗口并运行程序
	tidyUp()              //结束应用程序后执行tidyUp()函数
}

func tidyUp() {
	fmt.Println("Exited")
}
