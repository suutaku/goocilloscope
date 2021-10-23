/**
Connector for emulator, by default, this connector will output a sin wave
**/
package connector

import (
	"context"
	"math"
	"time"
)

type Emulator struct {
	ctx    context.Context
	cancel context.CancelFunc
	Connector
	buf     chan []float32
	waveBuf []float32
}

func NewEmulator(ctxf context.Context) *Emulator {
	ctx, cancel := context.WithCancel(ctxf)
	return &Emulator{
		ctx:    ctx,
		cancel: cancel,
		buf:    make(chan []float32),
	}
}

func (em *Emulator) GetBufferChannel() chan []float32 {
	return em.buf
}

func (em *Emulator) Open() error {
	go em.sinWave()
	go func() {
		tk := time.NewTicker(30 * time.Millisecond)
		for {
			select {
			case <-tk.C:
				em.buf <- em.waveBuf
				em.waveBuf = make([]float32, 0)
			case <-em.ctx.Done():
				tk.Stop()
				return
			}
		}
	}()
	return nil
}

func (em *Emulator) ReadBytes() ([]byte, error) {
	return nil, nil
}

func (em *Emulator) Close() {
	em.cancel()
}

func (em *Emulator) sinWave() {
	count := int64(0)
	for {
		s := math.Sin(float64(count) * math.Pi / 73)
		em.waveBuf = append(em.waveBuf, float32(s))
		//fmt.Println(s)
		time.Sleep(10 * time.Millisecond)
		count++
	}
}

func (em *Emulator) Name() string {
	return "Emulator"
}
