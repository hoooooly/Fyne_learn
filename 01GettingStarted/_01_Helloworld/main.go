/*
 * @Author: holy
 * @Date: 2022-04-01 00:59:02
 * @LastEditTime: 2022-04-01 01:02:11
 * @LastEditors: holy
 * @Description:
 * @FilePath: \Fyne_learn\01GettingStarted\_01_Helloworld\main.go
 */
package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()                            //创建应用程序
	myWindow := myApp.NewWindow("Hello")          //创建窗口
	myWindow.SetContent(widget.NewLabel("Hello")) //在窗口上设置控件

	myWindow.Show() //显示窗口
	myApp.Run()     //运行应用程序
	tidyUp()        //结束应用程序后执行tidyUp()函数
}

func tidyUp() {
	fmt.Println("Exited")
}
