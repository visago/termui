// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

// +build ignore

package main

import (
	"log"
	"math"

	ui "github.com/visago/termui/v3"
	"github.com/visago/termui/v3/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	sinData := func() [][]float64 {
		n := 220
		data := make([][]float64, 2)
		data[0] = make([]float64, n)
		data[1] = make([]float64, n)
		for i := 0; i < n; i++ {
			data[0][i] = 1 + math.Sin(float64(i)/5)
			data[1][i] = 1 + math.Cos(float64(i)/5)
		}
		return data
	}()

	p0 := widgets.NewPlot()
	p0.Title = "braille-mode Line Chart"
	p0.Data = sinData
	p0.DataLabels = []string{"00:00","00:01","00:02","00:03","00:04","00:05","00:06","00:07","00:08","00:09","00:10","00:11","00:12","00:13","00:14","00:15","00:16","00:17","00:18","00:19","00:20","00:21","00:22","00:23","00:24","00:25","00:26","00:27","00:28","00:29","00:30","00:31","00:32","00:33","00:34","00:35","00:36","00:37","00:38","00:39","00:40"}
	p0.SetRect(0, 0, 50, 15)
	p0.AxesColor = ui.ColorWhite
	p0.LineColors[0] = ui.ColorGreen
	p0.LabelAxesX = true

	p1 := widgets.NewPlot()
	p1.Title = "dot-mode line Chart"
	p1.Marker = widgets.MarkerDot
	p1.Data = [][]float64{[]float64{1, 2, 3, 4, 5}}
	p1.SetRect(50, 0, 75, 10)
	p1.DotMarkerRune = '+'
	p1.AxesColor = ui.ColorWhite
	p1.LineColors[0] = ui.ColorYellow
	p1.DrawDirection = widgets.DrawLeft
	p1.ShowAxesX = false

	p2 := widgets.NewPlot()
	p2.Title = "dot-mode Scatter Plot"
	p2.Marker = widgets.MarkerDot
	p2.Data = make([][]float64, 2)
	p2.Data[0] = []float64{1, 2, 3, 4, 5}
	p2.Data[1] = sinData[1][4:]
	p2.SetRect(0, 15, 50, 30)
	p2.AxesColor = ui.ColorWhite
	p2.LineColors[0] = ui.ColorCyan
	p2.PlotType = widgets.ScatterPlot

	p3 := widgets.NewPlot()
	p3.Title = "braille-mode Scatter Plot"
	p3.Data = make([][]float64, 2)
	p3.Data[0] = []float64{1, 2, 3, 4, 5}
	p3.Data[1] = sinData[1][4:]
	p3.SetRect(45, 15, 80, 30)
	p3.AxesColor = ui.ColorWhite
	p3.LineColors[0] = ui.ColorCyan
	p3.Marker = widgets.MarkerBraille
	p3.PlotType = widgets.ScatterPlot

	ui.Render(p0, p1, p2, p3)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}
