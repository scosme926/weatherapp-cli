package models

import (
	"time"
)

type TimeSeriesDatum struct {
	Temperature float64
	Humidity    float64
	Pressure    float64
	Co2         float64
	Tvoc        float64
	Timestamp   time.Time
}
