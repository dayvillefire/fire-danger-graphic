package main

import (
	"embed"
	"flag"
	"fmt"
	"image/color"
	"os"
	"time"

	"github.com/jbuchbinder/gg"
)

var (
	//go:embed resources/*
	embedfs embed.FS

	courtesyOf      = flag.String("courtesy-of", "Division of Forestry, Bureau of Natural Resources, Connecticut Department of Energy and Environmental Protection", "Courtesy attribution text")
	fireDangerLevel = flag.String("level", "HIGH", "LOW/MEDIUM/HIGH/EXTREME")
	outputFile      = flag.String("out", "danger.png", "Output file name")
)

func main() {
	flag.Parse()
	switch *fireDangerLevel {
	case "EXTREME":
		break
	case "HIGH":
		break
	case "MEDIUM":
		break
	case "LOW":
		break
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
	dc := gg.NewContext(1200, 630)

	// Background
	dc.SetHexColor("#000000")
	dc.DrawRectangle(0, 0, 1200.0, 630.0)
	dc.Fill()

	{
		var bgImage string
		switch *fireDangerLevel {
		case "EXTREME":
			bgImage = "resources/high.jpg"
		case "HIGH":
			bgImage = "resources/high.jpg"
		case "MEDIUM":
			bgImage = "resources/medium.jpg"
		case "LOW":
			bgImage = "resources/low.jpg"
		default:
			panic("Invalid fire danger level")
		}
		bg, err := gg.LoadJPGFS(embedfs, bgImage)
		if err != nil {
			panic(err)
		}
		//h := bg.Bounds().Size().Y
		dc.DrawImage(bg, 0, 0) //630-h)
	}

	{
		text := fmt.Sprintf("Fire Danger Level for %s is", time.Now().Format("Monday, Jan 2, 2006"))
		if err := dc.LoadFontFaceFS(embedfs, "resources/DejaVuSans-Bold.ttf", 40); err != nil {
			panic(err)
		}
		dc.SetHexColor("#000000")
		dc.DrawStringAnchored(text, 1200/2+4, 30+4, 0.5, 0.5)
		dc.SetHexColor("#FFFFFF")
		dc.DrawStringAnchored(text, 1200/2, 30, 0.5, 0.5)
	}

	{
		if err := dc.LoadFontFaceFS(embedfs, "resources/DejaVuSans-Bold.ttf", 128); err != nil {
			panic(err)
		}
		dc.SetHexColor("#FFFFFF")
		dc.DrawStringAnchored(*fireDangerLevel, 1200/2+5, 630/2+5, 0.5, 0.5)
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
		dc.DrawStringAnchored(*fireDangerLevel, 1200/2, 630/2, 0.5, 0.5)
	}

	// Background rounded rectangle behind courtesy text
	{
		dc.Push()
		dc.SetHexColor("#FFFFFF")
		dc.SetFillStyle(gg.NewSolidPattern(color.RGBA{128, 128, 128, 96}))
		dc.DrawRoundedRectangle(50, 590, 1100.00, 30.00, 10.00)
		dc.Fill()
		dc.Pop()
	}

	{
		if err := dc.LoadFontFaceFS(embedfs, "resources/DejaVuSans.ttf", 16); err != nil {
			panic(err)
		}
		dc.SetHexColor("#000000")
		dc.DrawStringAnchored("Courtesy of: "+*courtesyOf, 1200/2, 600, 0.5, 1)
	}

	err := dc.SavePNG(*outputFile)
	if err != nil {
		panic(err)
	}
}
