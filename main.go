package main

import (
	"flag"
	"fmt"
	"image/color"
	"time"

	"github.com/fogleman/gg"
)

var (
	fireDangerLevel = flag.String("level", "HIGH", "LOW/MEDIUM/HIGH/EXTREME")
	outputFile      = flag.String("out", "danger.png", "Output file name")
)

func main() {
	dc := gg.NewContext(1200, 630)

	// Background
	dc.SetHexColor("#000000")
	dc.DrawRectangle(0, 0, 1200.0, 630.0)
	dc.Fill()

	{
		bg, err := gg.LoadJPG("fire-bg.jpg")
		if err != nil {
			panic(err)
		}
		h := bg.Bounds().Size().Y
		dc.DrawImage(bg, 0, 630-h)
	}

	{
		text := fmt.Sprintf("Fire Danger Level for %s is", time.Now().Format("Monday, Jan 2, 2006"))
		dc.SetHexColor("#FFFFFF")
		if err := dc.LoadFontFace("DejaVuSans.ttf", 40); err != nil {
			panic(err)
		}
		dc.DrawStringAnchored(text, 1200/2, 30, 0.5, 0.5)
	}

	{
		switch *fireDangerLevel {
		case "EXTREME":
			dc.SetHexColor("#FF0000")
		case "HIGH":
			dc.SetHexColor("#FF0000")
		case "MEDIUM":
			dc.SetHexColor("#FF9900")
		case "LOW":
			dc.SetHexColor("#0000FF")
		}
		if err := dc.LoadFontFace("DejaVuSans-Bold.ttf", 128); err != nil {
			panic(err)
		}
		dc.SetHexColor("#FFFFFF")
		dc.DrawStringAnchored(*fireDangerLevel, 1200/2+5, 630/2+5, 0.5, 0.5)
		dc.SetHexColor("#FF0000")
		dc.DrawStringAnchored(*fireDangerLevel, 1200/2, 630/2, 0.5, 0.5)
	}

	{
		dc.Push()
		dc.SetHexColor("#FFFFFF")
		dc.SetFillStyle(gg.NewSolidPattern(color.RGBA{128, 128, 128, 96}))
		dc.DrawRoundedRectangle(50, 590, 1100.00, 30.00, 10.00)
		dc.Fill()
		dc.Pop()
	}

	{
		if err := dc.LoadFontFace("DejaVuSans.ttf", 16); err != nil {
			panic(err)
		}
		dc.SetHexColor("#000000")
		dc.DrawStringAnchored("Courtesy of: Division of Forestry, Bureau of Natural Resources, Connecticut Department of Energy and Environmental Protection", 1200/2, 600, 0.5, 1)
	}

	err := dc.SavePNG(*outputFile)
	if err != nil {
		panic(err)
	}
}
