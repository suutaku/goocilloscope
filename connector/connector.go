package connector

type Connector interface {
	Open() error
	Close()
	ReadBytes() ([]byte, error)
	GetBufferChannel() chan []float32
	Name() string
}
