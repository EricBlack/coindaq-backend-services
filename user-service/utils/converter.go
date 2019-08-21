package utils

import (
	"encoding/json"
	"time"
	"errors"
)

func Map2Json(input map[string]string) ([]byte, error) {
	output, err := json.Marshal(input)
	if err != nil {
		return output, err
	}
	return output, nil
}

func Json2Map(input []byte) (map[string]string, error) {
	var result map[string]string
	if err := json.Unmarshal(input, &result); err != nil {
		return result, err
	}
	return result, nil
}

//转换为UTC时间
func String2Time(timeString string) (time.Time, error) {
	format := "2006-01-02 15:04:05"
	timeInfo, err := time.Parse(format, timeString)
	if err != nil {
		return time.Time{}, errors.New("Time parameter invalid, correct time format is: '2006-01-02 15:04:05'")
	}

	return timeInfo, nil
}

//转换为本地时间
func String2TimeWithLocation(timeString string) (time.Time, error) {
	format := "2006-01-02 15:04:05"
	timeInfo, err := time.ParseInLocation(format, timeString, time.Local)
	if err != nil {
		return time.Time{}, errors.New("Time parameter invalid, correct time format is: '2006-01-02 15:04:05'")
	}

	return timeInfo, nil

}

func Time2String(timeInfo time.Time) string {
	format := "2006-01-02 15:04:05"
	timeString := timeInfo.Format(format)

	return timeString
}
