/*
 * @Author: holy
 * @Date: 2022-04-01 01:04:16
 * @LastEditTime: 2022-04-01 01:15:31
 * @LastEditors: holy
 * @Description:
 * @FilePath: \Fyne_learn\01GettingStarted\_02_UpdatingContent\main.go
 */
package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func updateTime2Label(l *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05") //设置时间
	l.SetText(formatted)                             //设置标签文本s
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

	// myWindow.Show() //显示窗口
	// myApp.Run()     //运行应用程序
	myWindow.ShowAndRun() //显示窗口并运行程序
	tidyUp()              //结束应用程序后执行tidyUp()函数
}

func tidyUp() {
	fmt.Println("Exited")
}
