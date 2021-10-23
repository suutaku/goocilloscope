/**
Connector for serial port
**/
package connector

import (
	"context"
	"fmt"
	"strconv"

	"go.bug.st/serial"
)

type Serial struct {
	Connector
	ctx          context.Context
	portName     string
	mode         *serial.Mode
	port         serial.Port
	buf          chan []float32
	quit         bool
	washCallback func([]byte) []float32
}

func NewSerial(ctxf context.Context, portName string, baudRate int) *Serial {

	mode := &serial.Mode{
		BaudRate: baudRate,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	return &Serial{
		portName:     portName,
		mode:         mode,
		ctx:          ctxf,
		buf:          make(chan []float32),
		washCallback: defaultWashCallback,
	}
}

func defaultWashCallback(data []byte) []float32 {
	tmp := make([]float32, 0)
	for _, v := range data {
		tmp64, _ := strconv.ParseFloat(string(v), 32)
		tmp = append(tmp, float32(tmp64))
	}
	return tmp
}

func (se *Serial) SetWashCallback(cb func([]byte) []float32) {
	se.washCallback = cb
}

func (se *Serial) Open() error {
	if se.portName == "" {
		ports, err := serial.GetPortsList()
		if err != nil || len(ports) < 1 {
			panic("can not found any valid serial port")
		}
		se.portName = ports[0]
		fmt.Println("no port name specific, using default: ", ports[0])
	}
	port, err := serial.Open(se.portName, se.mode)
	if err != nil {
		return err
	}
	se.port = port
	go func() {
		for !se.quit {
			tmp := make([]byte, 1024)
			n, _ := se.port.Read(tmp)
			se.buf <- se.washCallback(tmp[:n])
		}
	}()

	go func() {
		<-se.ctx.Done()
		se.quit = true

	}()
	return nil
}

func (se *Serial) Close() {
	se.port.Close()
}
func (se *Serial) ReadBytes() ([]byte, error) {
	return nil, nil
}
func (se *Serial) GetBufferChannel() chan []float32 {
	return se.buf
}

func (se *Serial) Name() string {
	return se.portName
}
