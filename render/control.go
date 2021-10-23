package render

import "fmt"

func (rd *Render) ZoomIn() {
	rd.zoomParam *= 0.8
	fmt.Println("zoom in: ", rd.zoomParam, rd.zoomOffsetStart)
}
func (rd *Render) ZoomOut() {
	rd.zoomParam /= 0.8
	if rd.zoomParam > 1 {
		rd.zoomParam = 1
	}
	rd.zoomOffsetStart = 0
	fmt.Println("zoom out: ", rd.zoomParam, rd.zoomOffsetStart)
}

func (rd *Render) Pause() {
	rd.pause = !rd.pause
	fmt.Println("pause")
}
