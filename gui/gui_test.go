package gui

import (
	"flag"
	"github.com/andlabs/ui"
	"github.com/woat/uimon"
	"log"
	"testing"
)

var f string

func init() {
	flag.StringVar(&f, "uimon", "", "use 'start' to run ui")
	flag.Parse()
}

func TestMain(m *testing.M) {
	switch f {
	case "run":
		log.Println("RUN FLAG")
		uimon.Start(Execute, func() {
			ui.QueueMain(ui.Quit)
		})
	case "start":
		log.Println("START FLAG")
		uimon.Starter()
	}
	select {}
}

//ooooooooo
