package render

import (
	"fmt"
	"image"
	"time"

	"github.com/eiannone/keyboard"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/size"
)

func (rd *Render) windowEnventHandler(src screen.Screen) {
	fmt.Println("window event handler open")
	for !rd.quit {
		switch e := rd.window.NextEvent().(type) {
		case size.Event:
			if e.WidthPx == 0 && e.HeightPx == 0 {
				fmt.Println("exit progress")
				rd.quit = true
				rd.cancel()
			} else {
				rd.ResizeWindow(e.WidthPx, e.HeightPx)
				//fmt.Println("fix size: [", rd.height, rd.width, "]")
				rd.windowBuffer, _ = src.NewBuffer(image.Point{X: rd.width, Y: rd.height})
			}
		case lifecycle.Event:
			if e.To == lifecycle.StageDead {
				fmt.Println("cancel done")
				rd.quit = true
				rd.cancel()
			}
		}
	}
}

func (rd *Render) keyboardEventHandler() {

	// keyboard event
	if err := keyboard.Open(); err != nil {
		panic(err)
	}

	fmt.Println("keyboard event handler open")
	for !rd.quit {
		char, key, err := keyboard.GetKey()
		if err != nil {
			continue
		}
		fmt.Println(key, err, char)
		switch key {
		case keyboard.KeyArrowUp:
			rd.ZoomIn()
		case keyboard.KeyArrowDown:
			rd.ZoomOut()
		case keyboard.KeyEsc:
			rd.quit = true
			rd.cancel()
		case keyboard.KeyCtrlC:
			rd.quit = true
			rd.cancel()
		case 32:
			rd.Pause()
		}
	}

}

func (rd *Render) dataEventHandler() {
	// read & draw process
	tk := time.NewTicker(1 * time.Second)
	fmt.Println("data event handler open")
	min, max := rd.normalizeValue()
	rd.draw(min, max)
	ch := rd.connector.GetBufferChannel()
	for !rd.quit {
		// pause
		if rd.pause {
			time.Sleep(100 * time.Millisecond)
			continue
		}
		select {
		case b := <-ch:
			for _, v := range b {
				rd.addBuffer(v)
			}
		case <-tk.C:
		}

		min, max := rd.normalizeValue()
		rd.draw(min, max)
	}

}
