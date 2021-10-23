package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	cli "github.com/jawher/mow.cli"
	"github.com/suutaku/goocilloscope/connector"
	myrender "github.com/suutaku/goocilloscope/render"
)

func washData(input []byte) []float32 {

	tmp := strings.Split(string(input), ":")
	res := make([]float32, 0)
	for i := 0; i < len(tmp); i++ {
		tmp64, _ := strconv.ParseFloat(string(tmp[i]), 32)
		res = append(res, float32(tmp64))
	}
	return res
}

func cmdEmulator(cmd *cli.Cmd) {
	ctx := context.Background()
	conn := connector.NewEmulator(ctx)
	rd := myrender.NewRender(ctx, 1280, 640, conn)
	rd.Start()
}

func cmdPortAudio(cmd *cli.Cmd) {
	ctx := context.Background()
	conn := connector.NewPortAudio(ctx)
	rd := myrender.NewRender(ctx, 1280, 640, conn)
	rd.Start()
}

func cmdSerial(cmd *cli.Cmd) {
	cmd.Spec = "PORT_NAME"
	pn := cmd.StringArg("PORT_NAME", "", "The serial port name")
	cmd.Spec = "BAUD_RATE"
	br := cmd.StringArg("BAUD_RATE", "", "The serial baud rate")
	if pn == nil {
		fmt.Println("PORT_NAME not set, using default value")
		//"/dev/cu.usbserial-12BP0136"
	}
	if br != nil {
		fmt.Println("BAUD_RATE not set, using default value: 9600")
		*br = "9600"
	}
	brn, err := strconv.ParseInt(*br, 10, 32)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	conn := connector.NewSerial(ctx, *pn, int(brn))
	rd := myrender.NewRender(ctx, 1280, 640, conn)
	rd.Start()
}

func main() {
	app := cli.App("goocilloscope", "A simple ocilloscope writen in Go")
	app.Command("source", "configure command", func(cmd *cli.Cmd) {
		cmd.Command("emulator", "A sine wave emulator", cmdEmulator)
		cmd.Command("serial", "A sine wave emulator", cmdSerial)
		cmd.Command("portaudio", "A sine wave emulator", cmdPortAudio)

	})
	app.Run(os.Args)
}
