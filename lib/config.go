package lib

import (
	"rpi-radio-alarm-desktop-client/lib/funcs"
	"rpi-radio-alarm-desktop-client/lib/hooks"

	"github.com/guark/guark/app"
	"github.com/guark/plugins/clipboard"
	"github.com/guark/plugins/dialog"
	"github.com/guark/plugins/notify"
)

// Funcs Exposed functions to guark Javascript api.
var Funcs = app.Funcs{
	"hello_world":  funcs.HelloWorld,
	"get_alarms":   funcs.GetAlarms,
	"get_alarm":    funcs.GetAlarm,
	"change_alarm": funcs.ChangeAlarm,
	"get_radio":    funcs.GetRadio,
	"start_radio":  funcs.StartRadio,
	"stop_radio":   funcs.StopRadio,
	"get_vals":     funcs.GetVals,
}

// Hooks App hooks.
var Hooks = app.Hooks{
	"created": hooks.Created,
	"mounted": hooks.Mounted,
}

// Plugins App plugins.
var Plugins = app.Plugins{
	"dialog":    &dialog.Plugin{},
	"notify":    &notify.Plugin{},
	"clipboard": &clipboard.Plugin{},
}
