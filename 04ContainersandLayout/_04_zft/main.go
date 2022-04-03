/*
 * @Author: holy
 * @Date: 2022-04-02 23:54:19
 * @LastEditTime: 2022-04-03 01:17:03
 * @LastEditors: holy
 * @Description:
 * @FilePath: \Fyne_learn\04ContainersandLayout\_04_zft\main.go
 */

package main

import (
	"image/color"
	"math/rand"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("zft")

	zft(w)

	w.Resize(fyne.NewSize(500, 500))
	w.ShowAndRun()
}

func zft(w fyne.Window) {
	c := container.NewWithoutLayout()
	width := 40
	max := 100
	data := genData(10, max)

	for i, d := range data {
		line := canvas.NewLine(rColor())
		line.Position1 = fyne.NewPos(float32(width*i+10), 100)
		line.Position2 = fyne.NewPos(float32(width*i+10), 100-d)
		line.StrokeWidth = 30
		text := canvas.NewText(strconv.FormatFloat(float64(d), 'f', 0, 64), color.Black)
		text.Move(line.Position2)
		text.TextSize = 10
		text.Alignment = fyne.TextAlignCenter
		c.Objects = append(c.Objects, line)
		c.Objects = append(c.Objects, text)
	}
	//折线图
	for i := 0; i < len(data)-1; i++ {
		line := canvas.NewLine(rColor())
		line.Position1 = fyne.NewPos(float32(width*i+10), 100-data[i])
		line.Position2 = fyne.NewPos(float32(width*(i+1)+10), 100-data[i+1])
		line.StrokeWidth = 1
		c.Objects = append(c.Objects, line)
	}
	w.SetContent(c)
}

// 产生随机颜色
func rColor() color.Color {
	return color.RGBA{uint8(rand.Intn(256)), uint8(rand.Intn(256)), uint8(rand.Intn(256)), 255}
}

// 产生随机数据
func genData(count, max int) (data []float32) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < count; i++ {
		data = append(data, float32(rand.Intn(max)))
	}
	return data
}
