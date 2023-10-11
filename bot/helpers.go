// Package bot contains helper functions for the bot.
package bot

// Importing necessary libraries.
import (
	"encoding/json"
	"math"

	"github.com/sirupsen/logrus"
)

// roundToNearestTick rounds the size to the nearest tick.
func roundToNearestTick(size, tick float64) *float64 {
	if tick <= 0 {
		return &size
	}
	numTicks := math.Round(size / tick)
	val := numTicks * tick
	return &val
}

// decodeAndPrintData decodes and prints data if show is true.
func decodeAndPrintData[T any](channel string, data []byte, show bool) *T {

	var t T
	err := json.Unmarshal(data, &t)
	if err != nil {
		logrus.Errorf("Failed to unmarshal data = %s", err.Error())
		return nil
	}

	if show {
		logrus.Infof("*** Channel = %s", channel)
		logrus.Info(string(data))
	}
	return &t

	// Here is one of the possible way to decode and then publish non-empty fields
	/*
		isTitle := false
		val := reflect.ValueOf(t)
		for i := 0; i < val.NumField(); i++ {
			valueField := val.Field(i)
			typeField := val.Type().Field(i)
			if !valueField.IsZero() {
				if !isTitle {
					isTitle = true
					logrus.Infof("Channel = %s", channel)
				}
				logrus.Infof("%s,\t = %v\n", typeField.Name, valueField.Interface())
			}
		}
	*/
}
