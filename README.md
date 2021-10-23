# goocilloscope
 A ocilloscope writen in GO.
 Supported serial input, portaudio input.


## Install

```bash
go get github.com/suutaku/goosilloscope
go build
```

## Usage
As a tool:
```bash
Usage: goocilloscope COMMAND [arg...]

A simple ocilloscope writen in Go
               
Commands:      
  source       need specific a signnal source
               
Run 'goocilloscope COMMAND --help' for more information on a command.
```
As a libray:

```go
	ctx := context.Background()
	conn := connector.NewPortAudio(ctx)
	rd := myrender.NewRender(ctx, 1280, 640, conn)
	rd.Start()
```

## Ohters
serial port default split recive bytes with `:`,you can set your costom data wash callback like this:

```go
func washData(input []byte) []float32 {

	tmp := strings.Split(string(input), ":")
	res := make([]float32, 0)
	for i := 0; i < len(tmp); i++ {
		tmp64, _ := strconv.ParseFloat(string(tmp[i]), 32)
		res = append(res, float32(tmp64))
	}
	return res
}
conn := connector.NewSerial(ctx, "/dev/usb-serial", 9600)
conn.SetWashCallback(washData)

```
For implamemt another input source, please see:

```bash
$REPO/connector/connector.go
```

