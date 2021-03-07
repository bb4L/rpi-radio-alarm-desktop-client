package funcs

import (
	"fmt"
	"strconv"

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

// GetAlarm gets the alarms
func GetAlarm(c app.Context) (interface{}, error) {
	logger := constants.GetLogger()
	helper := api.Helper{AlarmURL: Alarmurl, ExtraHeader: Extraheader, ExtreaHeaderValue: Extraheadervalue, Logger: logger}
	idx, err := strconv.ParseInt(fmt.Sprintf("%v", (c.Params["idx"])), 10, 32)
	if err != nil {
		return nil, err
	}
	return helper.GetAlarm(int(idx))
}
