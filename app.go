package main

import (
	"rpi-radio-alarm-desktop-client/lib"

	"github.com/guark/guark/app"
	"github.com/guark/guark/engine"
	"github.com/guark/guark/log"
)

func main() {

	a := &app.App{
		Log:   log.New("app"),
		Funcs: lib.Funcs,
		Embed: lib.Embeds,
	}

	if err := a.Use(engine.New(a)); err != nil {
		a.Log.Fatal(err)
	}

	defer a.Quit()

	if err := a.Run(); err != nil {
		a.Log.Fatal(err)
	}
}
