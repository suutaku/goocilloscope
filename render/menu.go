package render

import (
	"fmt"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

func (rd *Render) MakeStatusBar(dc *gg.Context, min, max float32) {
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic(err)
	}
	paddingX := (rd.height - rd.activeHeight) / 4
	// paddingY := (rd.height - rd.activeHeight) / 4
	face := truetype.NewFace(font, &truetype.Options{
		Size: float64(paddingX),
	})
	dc.SetFontFace(face)
	dc.SetRGBA(128.0, 64.0, 255.0, 128)
	anchorX := float64(paddingX)
	anchorY := float64(rd.activeHeight+(rd.height-rd.activeHeight)/2) + float64(paddingX)/2
	minText := fmt.Sprintf("Min: %f", min)
	w, _ := dc.MeasureString(minText)
	dc.DrawString(minText, anchorX, float64(anchorY))
	anchorX = anchorX + w + float64(paddingX)
	maxText := fmt.Sprintf("Max: %f", max)
	dc.DrawString(maxText, float64(anchorX), float64(anchorY))
	w, _ = dc.MeasureString(maxText)
	anchorX = anchorX + w + float64(paddingX)
	currentIdx := len(rd.buffer) - 1
	if currentIdx > -1 && currentIdx < len(rd.buffer) {
		currentText := fmt.Sprintf("Curr: %f", rd.buffer[currentIdx])
		dc.DrawString(currentText, float64(anchorX), float64(anchorY))
	}

	newPaddingX := float64(paddingX) / 2.0
	face = truetype.NewFace(font, &truetype.Options{
		Size: newPaddingX,
	})
	dc.SetFontFace(face)
	command1 := "↓ ZoomOut"
	w, h := dc.MeasureString(command1)
	anchorX = float64(rd.width) - float64(newPaddingX)*2 - w
	dc.DrawString(command1, float64(anchorX), float64(anchorY))
	dc.DrawRectangle(anchorX-float64(newPaddingX), anchorY-h-float64(newPaddingX), w+2*float64(newPaddingX), h+2*float64(newPaddingX))

	command1 = fmt.Sprintf("%.2f", 1/rd.zoomParam)
	w, h = dc.MeasureString(command1)
	anchorX = anchorX - float64(newPaddingX)*4 - w
	dc.DrawString(command1, float64(anchorX), float64(anchorY))
	dc.DrawRectangle(anchorX-float64(newPaddingX), anchorY-h-float64(newPaddingX), w+2*float64(newPaddingX), h+2*float64(newPaddingX))

	command1 = "↑ ZoomIn"
	w, h = dc.MeasureString(command1)
	anchorX = anchorX - float64(newPaddingX)*4 - w
	dc.DrawString(command1, float64(anchorX), float64(anchorY))
	dc.DrawRectangle(anchorX-float64(newPaddingX), anchorY-h-float64(newPaddingX), w+2*float64(newPaddingX), h+2*float64(newPaddingX))

	command1 = "space Pause"
	w, h = dc.MeasureString(command1)
	anchorX = anchorX - float64(newPaddingX)*4 - w
	dc.DrawString(command1, float64(anchorX), float64(anchorY))
	dc.DrawRectangle(anchorX-float64(newPaddingX), anchorY-h-float64(newPaddingX), w+2*float64(newPaddingX), h+2*float64(newPaddingX))
	dc.Stroke()
}

func (rd *Render) MakeBoarder(dc *gg.Context) {
	dc.SetRGBA(128.0, 64.0, 255.0, 32)
	dc.DrawRectangle(2.5, 2.5, float64(rd.width)-3.5, float64(rd.height)-3.5)
	dc.Stroke()
}

func (rd *Render) MakeInfoBar(dc *gg.Context) {

	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic(err)
	}
	paddingX := 5
	// paddingY := (rd.height - rd.activeHeight) / 4
	face := truetype.NewFace(font, &truetype.Options{
		Size: float64(paddingX) * 2.5,
	})
	dc.SetFontFace(face)
	command1 := "-> " + rd.connector.Name()
	w, h := dc.MeasureString(command1)
	recw := float64(w * 1.5)
	rech := float64(rd.activeHeight) / 12
	dc.SetRGBA(128.0, 64.0, 255.0, 32)

	dc.DrawRectangle(0, 0, recw, rech)
	dc.Fill()
	dc.Stroke()
	dc.SetRGBA(0.0, 0.0, 0.0, 32)

	dc.DrawString(command1, recw/2-w/2, rech/2+h/2)
	dc.Stroke()
}

func (rd *Render) MakeBaseLines(dc *gg.Context, min, max float32) {
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic(err)
	}
	paddingX := (rd.height - rd.activeHeight) / 4
	// paddingY := (rd.height - rd.activeHeight) / 4
	face := truetype.NewFace(font, &truetype.Options{
		Size: float64(paddingX),
	})
	dc.SetFontFace(face)
	dc.DrawString(fmt.Sprintf("%f", min), float64(paddingX), float64(rd.activeHeight/4)-float64(paddingX))
	dc.DrawString(fmt.Sprintf("%f", max), float64(paddingX), float64(rd.activeHeight*3/4)+float64(paddingX))
	// midle line
	dc.SetRGBA(128.0, 64.0, 255.0, 32)
	dc.DrawLine(0, float64(rd.activeHeight/2), float64(rd.activeWidth-1), float64(rd.activeHeight/2))
	dc.Stroke()
	// manu line
	dc.DrawLine(0, float64(rd.activeHeight), float64(rd.activeWidth-1), float64(rd.activeHeight))
	dc.Stroke()
	// positive half line
	dc.SetRGBA(128.0, 64.0, 255.0, 200)
	dc.DrawLine(0, float64(rd.activeHeight/4), float64(rd.activeWidth-1), float64(rd.activeHeight/4))

	// negative half line
	dc.DrawLine(0, float64(rd.activeHeight*3/4), float64(rd.activeWidth-1), float64(rd.activeHeight*3/4))
	dc.Stroke()

}
