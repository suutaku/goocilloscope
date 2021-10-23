package render

func (rd *Render) normalizeValue() (float32, float32) {
	max := float32(0)
	min := float32(0)
	for _, v := range rd.buffer {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}
	rd.normalBuffer = make([]float32, len(rd.buffer))
	for i := 0; i < len(rd.buffer); i++ {
		rd.normalBuffer[i] = rd.buffer[int(float32(i)*rd.zoomParam)] - min
		rd.normalBuffer[i] = rd.normalBuffer[i] / (max - min) // 0~1
		//fmt.Printf("normalize: %f", rd.normalBuffer[i])
		rd.normalBuffer[i] = rd.normalBuffer[i] + 0.5 //0.5~1.5
		rd.normalBuffer[i] = rd.normalBuffer[i] * float32(rd.activeHeight/2)

		//fmt.Printf(" %f\n", rd.normalBuffer[i])
	}
	return min, max
}

func (rd *Render) addBuffer(b float32) {
	rd.buffer = append(rd.buffer, float32(b))
	if len(rd.buffer) >= int(float32(rd.activeWidth)) {
		rd.buffer = rd.buffer[1:]
	}
}
