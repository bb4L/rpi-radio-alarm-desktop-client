package lib

import (
	"rpi-radio-alarm-desktop-client/lib/funcs"

	"github.com/guark/guark/app"
)

// Funcs Exposed functions to guark Javascript api.
var Funcs = app.Funcs{
	"get_alarms":   funcs.GetAlarms,
	"get_alarm":    funcs.GetAlarm,
	"change_alarm": funcs.ChangeAlarm,
	"create_alarm": funcs.CreateAlarm,
	"delete_alarm": funcs.DeleteAlarm,
	"get_radio":    funcs.GetRadio,
	"start_radio":  funcs.StartRadio,
	"stop_radio":   funcs.StopRadio,
	"get_vals":     funcs.GetVals,
}
