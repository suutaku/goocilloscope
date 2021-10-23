package test

import (
	"testing"

	"github.com/suutaku/goosilloscope/connector"
)

func TestAlsa(t *testing.T) {
	al := connector.NewAlsa()
	err := al.Open()
	if err != nil {
		panic(err)
	}
	al.Info()
}
