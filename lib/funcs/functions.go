package funcs

import (
	"fmt"
	// "reflect"
	"encoding/json"
	"strconv"

	"github.com/bb4L/rpi-radio-alarm-go-library/api"
	"github.com/bb4L/rpi-radio-alarm-go-library/constants"
	"github.com/bb4L/rpi-radio-alarm-go-library/types"
	"github.com/guark/guark/app"
)

func HelloWorld(c app.Context) (interface{}, error) {

	c.App.Log.Info("HelloWorld called")

	return nil, nil
}

var Alarmurl = "TEST_URL"
var Extraheader = "EXTRAHEADER"
var Extraheadervalue = "EXTRAHEADERVALUE"

func getHelper() api.Helper {
	logger := constants.GetLogger()
	return api.Helper{AlarmURL: Alarmurl, ExtraHeader: Extraheader, ExtreaHeaderValue: Extraheadervalue, Logger: logger}
}

// GetVals gets the vals
func GetVals(c app.Context) (interface{}, error) {
	return Alarmurl + " " + Extraheader + " " + Extraheadervalue, nil
}

// GetAlarms gets the alarms
func GetAlarms(c app.Context) (interface{}, error) {
	logger := constants.GetLogger()
	helper := api.Helper{AlarmURL: Alarmurl, ExtraHeader: Extraheader, ExtreaHeaderValue: Extraheadervalue, Logger: logger}
	alarms, err := helper.GetAlarms()
	if err != nil {
		return nil, err
	}
	data, jsonErr := json.Marshal(alarms)
	if jsonErr != nil {
		return jsonErr, nil
	}
	return string(data[:]), nil
}

// GetAlarm gets the alarms
func GetAlarm(c app.Context) (interface{}, error) {
	logger := constants.GetLogger()
	helper := api.Helper{AlarmURL: Alarmurl, ExtraHeader: Extraheader, ExtreaHeaderValue: Extraheadervalue, Logger: logger}
	idx, err := strconv.ParseInt(fmt.Sprintf("%v", (c.Params["idx"])), 10, 32)
	if err != nil {
		return nil, err
	}
	alarm, _ := helper.GetAlarm(int(idx))
	data, jsonErr := json.Marshal(alarm)
	if jsonErr != nil {
		return jsonErr, nil
	}
	return string(data[:]), nil
}

// ChangeAlarm change the alarm
func ChangeAlarm(c app.Context) (interface{}, error) {
	// paramAlarm := c.Params["alarm"].(types.Alarm)
	// paramAlarm, ok := c.Params["alarm"].(types.Alarm)
	// if !ok {
	// 	return "not ok", nil
	// }
	// c.App.Log.Debug(paramAlarm)
	// var alarm types.Alarm
	// alarmJson := make(types.Alarm)

	// alarm.Name = paramAlarm.Name
	// alarm.Days = paramAlarm[]
	// val := c.Params["alarm"]
	// val := c.Params.GetOr("alarm", types.Alarm{})
	// val, _ := c.GetOr("alarm", types.Alarm{}).(types.Alarm)
	val, ok := c.GetOr("alarm", "").(string)
	// err := json.Unmarshal([]byte(val), &alarm)
	if !ok {
		return "u", fmt.Errorf("get or err: " + val)
	}

	var alarm types.Alarm
	json.Unmarshal([]byte(val), &alarm)

	if len(val) < 4 {
		return "a", fmt.Errorf("val <5 : " + val)
	}

	// bytes := []byte(val)

	// return "val: " + val, nil

	// return alarm, nil

	idx, err := strconv.ParseInt(fmt.Sprintf("%v", (c.Params["idx"])), 10, 32)
	if err != nil {
		return "b", err
	}

	helper := getHelper()
	result_alarm, alarm_err := helper.ChangeAlarm(alarm, int(idx))
	if alarm_err != nil {
		return "c", alarm_err
	}

	// return result_alarm.Name + " idx: " + fmt.Sprint(int(idx)), nil

	data, jsonErr := json.Marshal(result_alarm)
	// data, jsonErr := json.Marshal(alarm)
	if jsonErr != nil {
		return "d", jsonErr
	}
	return string(data[:]), nil
}

// GetRadio gets the alarms
func GetRadio(c app.Context) (interface{}, error) {
	helper := getHelper()
	return helper.GetRadio()
}

// StartRadio starts the radio
func StartRadio(c app.Context) (interface{}, error) {
	helper := getHelper()
	return helper.StartRadio()
}

// StopRadio stops the radio
func StopRadio(c app.Context) (interface{}, error) {
	helper := getHelper()
	return helper.StopRadio()
}
