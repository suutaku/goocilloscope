/**
Connector for
**/
package connector

import (
	"context"
	"fmt"

	"github.com/gordonklaus/portaudio"
)

type PortAudio struct {
	Connector
	ctx    context.Context
	inBuf  chan []float32
	outBuf chan []float32
	host   *portaudio.HostApiInfo
	device *portaudio.StreamParameters
	*portaudio.Stream
	cancel context.CancelFunc
}

func NewPortAudio(ctx context.Context) *PortAudio {
	myCtx, cancel := context.WithCancel(ctx)
	e := &PortAudio{
		inBuf:  make(chan []float32, 1024),
		outBuf: make(chan []float32, 1024),
		ctx:    myCtx,
		cancel: cancel,
	}
	portaudio.Initialize()
	h, err := portaudio.DefaultHostApi()
	if err != nil {
		panic(err)
	}
	d := portaudio.LowLatencyParameters(h.DefaultInputDevice, h.DefaultOutputDevice)
	fmt.Println(d.Input.Device.Name)
	d.Input.Channels = 1
	d.Output.Channels = 1
	e.host = h
	e.device = &d
	return e
}

func (al *PortAudio) Open() error {
	go func() {
		s, err := portaudio.OpenStream(*al.device, al.processAudio)
		if err != nil {
			panic(err)
		}
		s.Start()
		al.Stream = s
		<-al.ctx.Done()
		al.Stream.Close()
		err = al.Stream.Stop()
		if err != nil {
			fmt.Println(err)
		}
		err = portaudio.Terminate()
		if err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("port audio open")
	return nil
}

func (al *PortAudio) Close() {

	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	fmt.Println("port audio close stream close")
	fmt.Println("port audio close terminate")
	al.cancel()
}

func (al *PortAudio) ReadBytes() ([]byte, error) {
	return nil, nil
}

func (al *PortAudio) GetBufferChannel() chan []float32 {
	return al.inBuf
}

func (al *PortAudio) GetOutPutBufferChannel() chan []float32 {
	return al.outBuf
}

func (al *PortAudio) Info() {

}

func (al *PortAudio) processAudio(in, out []float32) {
	//	fmt.Println("signal come")
	al.inBuf <- in
	//al.outBuf <- out
}

func (al *PortAudio) Name() string {
	return al.device.Input.Device.Name
}
