package funcs

import (
	"github.com/bb4L/rpi-radio-alarm-go-library/api"
	"github.com/bb4L/rpi-radio-alarm-go-library/constants"
	"github.com/guark/guark/app"
)

func HelloWorld(c app.Context) (interface{}, error) {

	c.App.Log.Info("HelloWorld called")

	return nil, nil
}

var Alarmurl = "TEST_URL"
var Extraheader = "EXTRAHEADER"
var Extraheadervalue = "EXTRAHEADERVALUE"

// GetVals gets the vals
func GetVals(c app.Context) (interface{}, error) {
	return Alarmurl + " " + Extraheader + " " + Extraheadervalue, nil
}

// GetAlarms gets the alarms
func GetAlarms(c app.Context) (interface{}, error) {
	logger := constants.GetLogger()
	helper := api.Helper{AlarmURL: Alarmurl, ExtraHeader: Extraheader, ExtreaHeaderValue: Extraheadervalue, Logger: logger}
	return helper.GetAlarms()
}
