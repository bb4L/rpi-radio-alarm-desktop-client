package funcs

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/bb4L/rpi-radio-alarm-go-library/api"
	"github.com/bb4L/rpi-radio-alarm-go-library/constants"
	"github.com/bb4L/rpi-radio-alarm-go-library/types"
	"github.com/guark/guark/app"
)

// Alarmurl url of the rpi-radio-alarm
var Alarmurl = "TEST_URL"

// Extraheader header for the requests
var Extraheader = "EXTRAHEADER"

// Extraheadervalue value for the extra header
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
	val, ok := c.GetOr("alarm", "").(string)
	if !ok {
		return nil, fmt.Errorf("get or err: " + val)
	}

	var alarm types.Alarm
	jsonErr := json.Unmarshal([]byte(val), &alarm)
	if jsonErr != nil {
		return nil, jsonErr
	}

	idx, err := strconv.ParseInt(fmt.Sprintf("%v", (c.Params["idx"])), 10, 32)
	if err != nil {
		return nil, err
	}

	helper := getHelper()
	resultAlarm, alarmErr := helper.ChangeAlarm(alarm, int(idx))
	if alarmErr != nil {
		return nil, alarmErr
	}

	data, jsonErr := json.Marshal(resultAlarm)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return string(data[:]), nil
}

// CreateAlarm creates a new alarm
func CreateAlarm(c app.Context) (interface{}, error) {
	val, ok := c.GetOr("alarm", "").(string)
	if !ok {
		return nil, fmt.Errorf("get or err: " + val)
	}

	var alarm types.Alarm
	err := json.Unmarshal([]byte(val), &alarm)
	if err != nil {
		return nil, err
	}

	helper := getHelper()
	alarms, addErr := helper.AddAlarm(alarm)

	if addErr != nil {
		return nil, addErr
	}

	data, jsonErr := json.Marshal(alarms)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return string(data[:]), nil
}

// DeleteAlarm deletes the alarm on the passed in index
func DeleteAlarm(c app.Context) (interface{}, error) {
	idx, err := strconv.ParseInt(fmt.Sprintf("%v", (c.Params["idx"])), 10, 32)
	if err != nil {
		return nil, err
	}
	helper := getHelper()
	alarms, err := helper.DeleteAlarm(int(idx))
	if err != nil {
		return nil, err
	}

	data, jsonErr := json.Marshal(alarms)
	if jsonErr != nil {
		return nil, jsonErr
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
